package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"time"
	"tupike_hotel/pkg/config"
	"tupike_hotel/pkg/types"

	// "github.com/Tech-Kenya/africastalking-sms-lib"
	"github.com/golang-jwt/jwt/v5"
	"github.com/tech-kenya/africastalking-sms-lib"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateNewCustomer(ctx context.Context, user *types.Customer, otp string) error {
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPass
	err = s.repo.InsertUnverified(ctx, user, otp)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LoginCustomer(ctx context.Context, email, password string) (string, error) {
	userFound, err := s.repo.LookUpCustomer(context.Background(), email)
	if err != nil || userFound == nil {
		return "", fmt.Errorf("user not found on the database: %v", err)
	}
	err = checkPass(userFound.Password, password)
	if err != nil {
		return "", errors.New("wrong password")
	}
	// proceed to assigning a token
	accessToken, err := createToken(userFound)
	if err != nil {
		return "", fmt.Errorf("error while generating access token: %v", err)
	}

	return accessToken, nil
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

// will add expiration later
func (s *Service) GenerateOTP() string {
	return strconv.Itoa(rand.Intn(9000) + 1000)
}

func (s *Service) SendSMS(phoneNumber, Otp string) (any, error) {
	client, err := africastalking.NewSMSClient(config.Envs.AtApiKey, config.Envs.AtUserName,
		config.Envs.AtShortCode, config.Envs.AtEnvironment)
	log.Println(client)
	if err != nil || client == nil {
		return "", err
	}
	resp, err := client.SendSMS(phoneNumber, Otp)
	if err != nil {
		return "", err
	}
	return resp, nil
}

func (s *Service) SendEmail(email, Otp string) error {
	return nil
}

func createToken(user *types.Customer) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": user.IsAdmin,
		"sub":  user.Email,
		"iss":  "todoApp",
		"exp":  time.Now().Add(time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})
	log.Println(token)
	return token.SignedString([]byte(config.Envs.SecretKey))
}

// todo
func validateToken() (bool, error) {
	return false, nil
}
