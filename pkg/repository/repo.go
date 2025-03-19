// interact with the database
package repository

import (
	"context"
	"tupike_hotel/pkg/types"

	"github.com/uptrace/bun"
)

type Repository struct {
	db *bun.DB
}

type RepoInterface interface {
	InsertCustomer(ctx context.Context, user *types.Customer) error
	LookUpCustomer(ctx context.Context, email string) (*types.Customer, error)
}

func NewRepository(db *bun.DB) RepoInterface {
	return &Repository{
		db: db,
	}
}
