package services

import (
	"context"
	"tupike_hotel/pkg/repository"
	"tupike_hotel/pkg/types"

	"github.com/go-playground/validator"
)

type Service struct {
	repo      repository.RepoInterface
	Validator *validator.Validate
}

type ServiceInterface interface {
	Validate(i any) error
	GetValidationErrors(err error) map[string]string
	CreateNewCustomer(ctx context.Context, user *types.Customer, otp string) error
	LoginCustomer(ctx context.Context, email, password string) (string, error)
	GenerateOTP() string
	SendSMS(phoneNumber, Otp string) (any, error)
	SendEmail(email, Otp string) error
}

func NewService(repo repository.RepoInterface, validator *validator.Validate) ServiceInterface {
	return &Service{
		repo:      repo,
		Validator: validator,
	}
}
