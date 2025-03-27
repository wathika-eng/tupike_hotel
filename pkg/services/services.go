package services

import (
	"tupike_hotel/pkg/repository"

	"github.com/go-playground/validator"
)

type Service struct {
	customerRepo *repository.CustomerRepo
	foodRepo     *repository.FoodRepo
	orderRepo    *repository.OrdersRepo
	validator    *validator.Validate
}

func NewService(customerRepo *repository.CustomerRepo, foodRepo *repository.FoodRepo,
	orderRepo *repository.OrdersRepo, validator *validator.Validate) *Service {
	return &Service{
		customerRepo: customerRepo,
		foodRepo:     foodRepo,
		orderRepo:    orderRepo,
		validator:    validator,
	}
}

// type AuthServices interface {
// 	GetValidationErrors(err error) map[string]string
// 	CreateNewCustomer(ctx context.Context, user *types.Customer) error
// 	LoginCustomer(ctx context.Context, email, password string) (string, error)
// 	Validate(i any) error
// }

// type OTPServices interface {
// 	GenerateOTP() string
// 	SendOTP(ctx context.Context, phoneNumber, email, Otp string) (any, error)
// 	CheckOTP(ctx context.Context, email, otp string) error
// }
