// interact with the database
package repository

import (
	"context"
	"tupike_hotel/pkg/types"

	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

type Repository struct {
	db    *bun.DB
	redis *redis.Client
}

type RepoInterface interface {
	InsertUnverified(ctx context.Context, user *types.Customer, otp string) error
	InsertCustomer(ctx context.Context, user *types.Customer) error
	LookUpCustomer(ctx context.Context, email string) (*types.Customer, error)
	InsertFood(ctx context.Context, food *types.FoodItem) error
	InsertOrder(ctx context.Context, order *types.Order) error
	Cleanup(ctx context.Context, user *types.Customer) error
}

func NewRepository(db *bun.DB, redis *redis.Client) RepoInterface {
	return &Repository{
		db:    db,
		redis: redis,
	}
}
