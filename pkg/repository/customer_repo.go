package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tupike_hotel/pkg/types"

	"github.com/jackc/pgx/v5/pgconn"
)

// create a new user
func (r Repository) InsertCustomer(ctx context.Context, user *types.Customer) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
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
func (r Repository) LookUpCustomer(ctx context.Context, email string) (*types.Customer, error) {
	var user types.Customer
	err := r.db.NewSelect().Model(&user).Where("email = ?", email).Limit(1).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}
