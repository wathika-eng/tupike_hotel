package repository

import (
	"context"
	"errors"
	"fmt"
	"tupike_hotel/pkg/types"

	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

type RepoInterface interface {
	InsertCustomer(ctx context.Context, user *types.Customer) error
}

func NewRepository(db *bun.DB) RepoInterface {
	return &Repository{
		db: db,
	}
}

func (r Repository) InsertCustomer(ctx context.Context, user *types.Customer) error {
	_, err := r.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return errors.New(fmt.Sprintf("error inserting new user to the database: %v", err.Error()))
	}
	return nil
}
