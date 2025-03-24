package repository

import (
	"context"
	"fmt"
	"tupike_hotel/pkg/types"

	"github.com/jackc/pgx/v5/pgconn"
)

func (r Repository) InsertFood(ctx context.Context, food *types.FoodItem) error {
	_, err := r.db.NewInsert().Model(food).Exec(ctx)
	if err != nil {
		// Check for unique constraint violation (DB-specific handling)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // 23505 = unique_violation
			return fmt.Errorf("food %s already exists", food.Item)
		}
		return fmt.Errorf("error inserting new food item into the database: %w", err)
	}
	return nil
}
