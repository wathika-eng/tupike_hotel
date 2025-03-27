// interact with the database
package repository

import (
	"github.com/redis/go-redis/v9"
	"github.com/uptrace/bun"
)

type Repository struct {
	CustomerRepo *CustomerRepo
	FoodRepo     *FoodRepo
	OrderRepo    *OrdersRepo
}

type DatabaseManager struct {
	DB    *bun.DB
	Redis *redis.Client
}

func NewDatabaseManager(db *bun.DB, redis *redis.Client) *DatabaseManager {
	return &DatabaseManager{
		DB:    db,
		Redis: redis,
	}
}

func NewRepository(db *DatabaseManager) *Repository {
	return &Repository{
		CustomerRepo: NewCustomerRepo(db),
		FoodRepo:     NewFoodRepo(db),
		OrderRepo:    NewOrdersRepo(db),
	}
}

// type AuthRepo interface {
// 	CheckOTP(ctx context.Context, email, otp string) error
// }

// type FoodRepo interface {
// 	InsertFood(ctx context.Context, food *types.FoodItem) error
// 	GetFood(ctx context.Context) error
// }

// type OrderRepo interface {
// 	InsertOrder(ctx context.Context, order *types.Order) error
// }
