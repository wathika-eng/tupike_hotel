package services

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"tupike_hotel/pkg/config"

	"github.com/tech-kenya/africastalking-sms-lib"
)

// will add expiration later
func (s *Service) GenerateOTP() string {
	return strconv.Itoa(rand.Intn(9000) + 1000)
}

func (s *Service) SendOTP(phoneNumber, email, Otp string) (any, error) {
	client, err := africastalking.NewSMSClient(config.Envs.AtApiKey, config.Envs.AtUserName,
		config.Envs.AtShortCode, config.Envs.AtEnvironment)

	if err != nil || client == nil {
		return "", err
	}
	// resp, err := client.SendSMS(phoneNumber, Otp)
	// if err != nil {
	// 	return "", err
	// }
	// implement email later
	return "", nil
}

func (s *Service) CheckOTP(ctx context.Context, email, otp string) error {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(otp) == "" {
		return errors.New("email or OTP cannot be blank")
	}
	return s.repo.CheckOTP(ctx, email, otp)
}
