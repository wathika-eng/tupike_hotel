package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tupike_hotel/pkg/types"

	"github.com/google/uuid"
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

// inserts food
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

// fetches food
func (r *FoodRepo) GetFood(ctx context.Context) ([]types.FoodItem, error) {
	var foods []types.FoodItem
	err := r.db.DB.NewSelect().Model(&foods).Scan(ctx)
	if err != nil {
		return nil, fmt.Errorf("error fetching food items: %w", err)
	}
	return foods, nil
}

func (r *FoodRepo) LookupFood(ctx context.Context, identifier string) (*types.FoodItem, error) {
	var food types.FoodItem

	query := r.db.DB.NewSelect().Model(&food).Limit(1)

	// Check if identifier is a valid UUID
	if _, err := uuid.Parse(identifier); err == nil {
		query.Where("id = ?", identifier)
	} else {
		query.Where("item ILIKE ?", identifier)
	}

	err := query.Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("food not found")
		}
		return nil, err
	}

	return &food, nil
}

func (r *FoodRepo) UpdateFood(ctx context.Context,
	food *types.FoodItem, orderedQuantity int) error {
	_, err := r.db.DB.NewUpdate().Model(food).
		Set("quantity = quantity - ?", orderedQuantity).
		Set("order_freq = order_freq + ?", orderedQuantity).
		Where("id = ? AND quantity >= ?", food.ID, orderedQuantity).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("unable to update food data")
	}
	return nil
}
