package services

import (
	"context"
	"errors"
	"fmt"
	"tupike_hotel/pkg/types"
)

func (s *Service) PlaceOrder(ctx context.Context, order *types.Order) error {
	food, err := s.foodRepo.LookupFood(ctx, order.FoodID.String())
	if err != nil {
		return err
	}
	if food.Quantity <= 0 {
		return errors.New("fooditem is not currently available")
	}
	order.AmountTotal = float64((float64(order.Quantity) * food.Price) - order.Discount)

	if order.AmountTotal < 0 {
		return errors.New("total amount cannot be negative")
	}
	// food.Quantity -= order.Quantity
	// food.OrderFreq += order.Quantity
	if err := s.foodRepo.UpdateFood(ctx, food, order.Quantity); err != nil {
		return err
	}
	if err := s.orderRepo.InsertOrder(ctx, order); err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}
	return nil
}
