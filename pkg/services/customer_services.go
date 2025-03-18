package services

import (
	"context"
	"errors"
	"tupike_hotel/pkg/repository"
	"tupike_hotel/pkg/types"

	"github.com/go-playground/validator"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      repository.RepoInterface
	Validator *validator.Validate
}

type ServiceInterface interface {
	CreateNewCustomer(ctx context.Context, user *types.Customer) error
	LoginCustomer(ctx context.Context, email string) error
	Validate(i any) error
	GetValidationErrors(err error) map[string]string
}

func NewService(repo repository.RepoInterface, validator *validator.Validate) ServiceInterface {
	return &Service{
		repo:      repo,
		Validator: validator,
	}
}

func (s *Service) CreateNewCustomer(ctx context.Context, user *types.Customer) error {
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPass
	err = s.repo.InsertCustomer(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LoginCustomer(ctx context.Context, email string) error {
	err := s.repo.LookUpCustomer(context.Background(), email)
	if err != nil {
		return errors.New("user not found on the database")
	}
	return nil
}
func hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("could not hash the password")
	}
	return string(hashedPass), nil
}
