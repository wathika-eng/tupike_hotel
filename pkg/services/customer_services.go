package services

import (
	"context"
	"errors"
	"tupike_hotel/pkg/repository"
	"tupike_hotel/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

type CustomerService struct {
	repo repository.RepoInterface
}

type CustomerServiceInterface interface {
	CreateNewCustomer(ctx context.Context, user *types.Customer) error
}

func NewCustomerService(repo repository.RepoInterface) CustomerServiceInterface {
	return &CustomerService{
		repo: repo,
	}
}

func (s *CustomerService) CreateNewCustomer(ctx context.Context, user *types.Customer) error {
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

func hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("could not hash the password")
	}
	return string(hashedPass), nil
}
