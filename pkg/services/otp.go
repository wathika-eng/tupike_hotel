package services

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"strings"
	"tupike_hotel/pkg/config"

	"github.com/resend/resend-go/v2"
	africastalking "github.com/tech-kenya/africastalkingsms"
)

// will add expiration later
func (s *Service) GenerateOTP() string {
	return strconv.Itoa(rand.Intn(9000) + 1000)
}

func (s *Service) SendOTP(ctx context.Context, phoneNumber, email, Otp string) (any, error) {
	// sms logic
	client, err := africastalking.NewSMSClient(config.Envs.AtApiKey, config.Envs.AtUserName,
		config.Envs.AtShortCode, config.Envs.AtEnvironment)

	if err != nil || client == nil {
		return "", err
	}
	// resp, err := client.SendSMS(phoneNumber, Otp)
	// if err != nil {
	// 	return "", err
	// }

	// email logic
	_ = resend.NewClient(config.Envs.RESEND_API_KEY)
	// params := &resend.SendEmailRequest{
	// 	From:    "Tupike Hotel <onboarding@resend.dev>",
	// 	To:      []string{email},
	// 	Html:    fmt.Sprintf("<strong>OTP Code:%v</strong>", Otp),
	// 	Subject: "Verification Code",
	// 	// Cc:      []string{"cc@example.com"},
	// 	// Bcc:     []string{"bcc@example.com"},
	// 	ReplyTo: "replyto@example.com",
	// }

	// sent, err := emailClient.Emails.SendWithContext(ctx, params)
	// if err != nil {
	// 	return "", err
	// }

	return "", nil
}

func (s *Service) CheckOTP(ctx context.Context, email, otp string) error {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(otp) == "" {
		return errors.New("email or OTP cannot be blank")
	}
	return s.customerRepo.CheckOTP(ctx, email, otp)
}
