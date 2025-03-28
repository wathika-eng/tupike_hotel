package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"time"
	"tupike_hotel/pkg/types"

	"github.com/google/uuid"
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
// 	LookUpCustomer(ctx context.Context, customerID string) (*types.Customer, error)
// 	Cleanup(ctx context.Context, user *types.Customer) error
// 	UpdateLoginTime(ctx context.Context, customerID string) error
// }

// create a new user
func (r *CustomerRepo) InsertCustomer(ctx context.Context, user *types.Customer) error {
	_, err := r.db.DB.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		// Check for unique constraint violation (DB-specific handling)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // 23505 = unique_violation
			return fmt.Errorf("user with customerID %s already exists", user.Email)
		}
		return fmt.Errorf("error inserting new user into the database: %w", err)
	}
	return nil
}

// check if a user exists in the database with customerID
func (r *CustomerRepo) LookUpCustomer(ctx context.Context, id string) (*types.Customer, error) {
	var user types.Customer

	_, err := uuid.Parse(id)
	if err == nil {
		err = r.db.DB.NewSelect().
			Model(&user).
			Where("id = ?", id).
			Limit(1).
			Scan(ctx)
	} else if isEmail(id) {
		err = r.db.DB.NewSelect().
			Model(&user).
			Where("email = ?", id).
			Limit(1).
			Scan(ctx)
	} else {
		return nil, errors.New("invalid ID format: not uuid or email")
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *CustomerRepo) CheckOTP(ctx context.Context, customerID, otp string) error {
	// check if user is in the database
	customer, err := r.LookUpCustomer(ctx, customerID)
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
	_, err = r.db.DB.NewUpdate().Model(customer).Column("verified").Where("id = ?", customerID).Exec(ctx)
	if err != nil {
		return err
	}
	customer.OTP = "0"
	_, err = r.db.DB.NewUpdate().Model(customer).Column("otp").Where("customerID = ?", customerID).Exec(ctx)
	return err
}

func (r *CustomerRepo) UpdateLoginTime(ctx context.Context, customerID string) error {
	customer, _ := r.LookUpCustomer(ctx, customerID)
	customer.LastLogin = time.Now()
	_, err := r.db.DB.NewUpdate().Model(customer).Column("last_login").Where("customerID = ?", customerID).Exec(ctx)
	return err
}

// cleanup deletes users who are unverified if 7days are over
func (r *CustomerRepo) Cleanup(ctx context.Context, user *types.Customer) error {
	_, err := r.db.DB.NewDelete().Model(user).Where("created_at < now() - interval '7 days'").Exec(ctx)
	return err
}

func isEmail(s string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(s)
}
