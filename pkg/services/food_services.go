package services

import (
	"context"
	"tupike_hotel/pkg/types"
)

func (s *Service) AddFood(ctx context.Context, food *types.FoodItem) error {
	return s.foodRepo.InsertFood(ctx, food)
}

func (s *Service) FetchFood(ctx context.Context) ([]types.FoodItem, error) {
	food, err := s.foodRepo.GetFood(ctx)
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (s *Service) CheckFood(ctx context.Context, foodName string) (*types.FoodItem, error) {
	foodData, err := s.foodRepo.LookupFood(ctx, foodName)
	if err != nil {
		return nil, err
	}
	return foodData, nil
}
