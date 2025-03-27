package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
	"tupike_hotel/pkg/types"

	"github.com/jackc/pgx/v5/pgconn"
)

type CustomerRepo struct {
	db *DatabaseManager
}

func NewCustomerRepo(db *DatabaseManager) *CustomerRepo {
	return &CustomerRepo{
		db: db,
	}
}

// type Customer interface {
// 	InsertCustomer(ctx context.Context, user *types.Customer) error
// 	LookUpCustomer(ctx context.Context, email string) (*types.Customer, error)
// 	Cleanup(ctx context.Context, user *types.Customer) error
// 	UpdateLoginTime(ctx context.Context, email string) error
// }

// create a new user
func (r *CustomerRepo) InsertCustomer(ctx context.Context, user *types.Customer) error {
	_, err := r.db.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		// Check for unique constraint violation (DB-specific handling)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // 23505 = unique_violation
			return fmt.Errorf("user with email %s already exists", user.Email)
		}
		return fmt.Errorf("error inserting new user into the database: %w", err)
	}
	return nil
}

// check if a user exists in the database with email
func (r *CustomerRepo) LookUpCustomer(ctx context.Context, email string) (*types.Customer, error) {
	var user types.Customer
	err := r.db.DB.NewSelect().Model(&user).Where("email = ?", email).Limit(1).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *CustomerRepo) CheckOTP(ctx context.Context, email, otp string) error {
	// check if user is in the database
	customer, err := r.LookUpCustomer(ctx, email)
	// if not in database, return
	if err != nil {
		return err
	}
	// is user is in the database but verified = true, return
	if customer.Verified {
		return errors.New("user is already verified")
	}
	// if user request body OTP is not same as OTP in the database, return
	if customer.OTP != otp {
		return errors.New("wrong otp code")
	}
	// change user to verified = true and set OTP = 0 (or drop the column) in the database
	customer.Verified = true
	_, err = r.db.DB.NewUpdate().Model(customer).Column("verified").Where("email = ?", email).Exec(ctx)
	if err != nil {
		return err
	}
	customer.OTP = "0"
	_, err = r.db.DB.NewUpdate().Model(customer).Column("otp").Where("email = ?", email).Exec(ctx)
	return err
}

func (r *CustomerRepo) UpdateLoginTime(ctx context.Context, email string) error {
	customer, _ := r.LookUpCustomer(ctx, email)
	customer.LastLogin = time.Now()
	_, err := r.db.DB.NewUpdate().Model(customer).Column("last_login").Where("email = ?", email).Exec(ctx)
	return err
}

// cleanup deletes users who are unverified if 7days are over
func (r *CustomerRepo) Cleanup(ctx context.Context, user *types.Customer) error {
	_, err := r.db.DB.NewDelete().Model(user).Where("created_at < now() - interval '7 days'").Exec(ctx)
	return err
}
