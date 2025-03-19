package services

import (
	"context"
	"errors"
	"fmt"
	"tupike_hotel/pkg/types"

	"golang.org/x/crypto/bcrypt"
)

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

func (s *Service) LoginCustomer(ctx context.Context, email, password string) error {
	userFound, err := s.repo.LookUpCustomer(context.Background(), email)
	if err != nil || userFound == nil {
		return fmt.Errorf("user not found on the database: %v", err)
	}
	err = checkPass(userFound.Password, password)
	if err != nil {
		return errors.New("wrong password")
	}
	// proceed to assigning a token
	return nil
}

// hashes the password before saving to the database
func hashPassword(password string) (string, error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("could not hash the password")
	}
	return string(hashedPass), nil
}

// check if the supplied password matches the hashed password on the db
func checkPass(hashedPass, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
}
