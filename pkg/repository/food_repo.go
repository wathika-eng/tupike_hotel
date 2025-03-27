package repository

import (
	"context"
	"fmt"
	"tupike_hotel/pkg/types"

	"github.com/jackc/pgx/v5/pgconn"
)

type FoodRepo struct {
	db *DatabaseManager
}

func NewFoodRepo(db *DatabaseManager) *FoodRepo {
	return &FoodRepo{
		db: db,
	}
}
func (r *FoodRepo) InsertFood(ctx context.Context, food *types.FoodItem) error {
	_, err := r.db.DB.NewInsert().Model(food).Exec(ctx)
	if err != nil {
		// Check for unique constraint violation (DB-specific handling)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // 23505 = unique_violation
			return fmt.Errorf("food %s already exists", food.Item)
		}
		return fmt.Errorf("error inserting new food item into the database: %w", err)
	}
	return nil
}

func (r *FoodRepo) GetFood(ctx context.Context) ([]types.FoodItem, error) {
	var foods []types.FoodItem
	err := r.db.DB.NewSelect().Model(&foods).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching food items: %w", err)
	}
	return foods, nil
}
