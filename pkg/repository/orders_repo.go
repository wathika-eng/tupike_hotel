package repository

import (
	"context"
	"fmt"
	"tupike_hotel/pkg/types"

	"github.com/jackc/pgx/v5/pgconn"
)

type OrdersRepo struct {
	db *DatabaseManager
}

func NewOrdersRepo(db *DatabaseManager) *OrdersRepo {
	return &OrdersRepo{
		db: db,
	}
}

func (r *OrdersRepo) InsertOrder(ctx context.Context, order *types.Order) error {
	_, err := r.db.DB.NewInsert().Model(order).Exec(ctx)
	if err != nil {
		// Check for unique constraint violation (DB-specific handling)
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" { // 23505 = unique_violation
			return fmt.Errorf("order %s already exists", order.ID)
		}
		return fmt.Errorf("error inserting new order item into the database: %w", err)
	}
	return nil
}
