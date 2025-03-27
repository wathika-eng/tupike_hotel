package services

import (
	"context"
	"tupike_hotel/pkg/types"
)

func (s *Service) FetchFood(ctx context.Context) ([]types.FoodItem, error) {
	food, err := s.foodRepo.GetFood(ctx)
	if err != nil {
		return nil, err
	}
	return food, nil
}
