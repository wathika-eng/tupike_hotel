package services

import (
	"context"
	"errors"
	"fmt"
	"tupike_hotel/pkg/types"
)

func (s *Service) PlaceOrder(ctx context.Context, order *types.Order) error {
	food, err := s.foodRepo.LookupFood(ctx, order.FoodItem)
	if err != nil {
		return err
	}
	if food.Quantity <= 0 {
		return errors.New("fooditem is not currently available")
	}
	if err := s.orderRepo.InsertOrder(ctx, order); err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}
	return nil
}
