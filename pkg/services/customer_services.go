package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"time"
	"tupike_hotel/pkg/config"
	"tupike_hotel/pkg/types"

	// "github.com/Tech-Kenya/africastalking-sms-lib"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateNewCustomer(ctx context.Context, user *types.Customer) error {
	hashedPass, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPass
	err = s.customerRepo.InsertCustomer(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LoginCustomer(ctx context.Context, email, password string) (string, error) {
	userFound, err := s.customerRepo.LookUpCustomer(context.Background(), email)
	if err != nil || userFound == nil {
		return "", fmt.Errorf("user not found on the database: %v", err)
	}

	err = checkPass(userFound.Password, password)
	if err != nil {
		return "", errors.New("wrong password")
	}
	// if userFound.Verified == false {
	// 	return "", errors.New("user not verified")
	// }
	// proceed to assigning a token
	accessToken, err := CreateToken(userFound)
	if err != nil {
		return "", fmt.Errorf("error while generating access token: %v", err)
	}
	s.customerRepo.UpdateLoginTime(ctx, email)
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

func CreateToken(user *types.Customer) (string, error) {
	claims := jwt.MapClaims{
		"admin": user.IsAdmin,
		"sub":   user.ID,
		"iss":   "todoApp",
		"exp":   time.Now().Add(7 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.Envs.SecretKey))
	if err != nil {
		log.Println("JWT signing error:", err)
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Envs.SecretKey), nil
	})

	if err != nil {
		log.Println("Token verification failed:", err)
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
