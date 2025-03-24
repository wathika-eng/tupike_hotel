package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"
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

func (r Repository) InsertUnverified(ctx context.Context, user *types.Customer, otp string) error {
	key := fmt.Sprintf("unverified_user:%s", user.Email)
	exists, err := r.redis.Exists(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to check email existence: %v", err)
	}
	if exists > 0 {
		return fmt.Errorf("email %s already exists", user.Email)
	}

	data, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user data: %v", err)
	}

	if err := r.redis.Set(ctx, key, data, 24*time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to store unverified user: %v", err)
	}

	return nil
}

// cleanup deletes users who are unverified if 7days are over
func (r Repository) Cleanup(ctx context.Context, user *types.Customer) error {
	_, err := r.db.NewDelete().Model(user).Where("created_at < now() - interval '7 days'").Exec(ctx)
	return err
}
