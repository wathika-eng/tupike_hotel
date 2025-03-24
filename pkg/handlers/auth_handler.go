package handlers

import (
	"context"
	"net/http"
	"strings"
	resp "tupike_hotel/pkg/response"
	"tupike_hotel/pkg/types"

	"github.com/labstack/echo/v4"
)

func (h *CustomerHandler) CreateUser(c echo.Context) error {
	var request struct {
		customer          *types.Customer
		OTPDeliveryMethod string `json:"otp_delivery_method" validate:"required,oneof=sms email"`
	}

	err := c.Bind(&request)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}
	err = h.service.Validate(request.customer)
	if err != nil {
		validationErrors := h.service.GetValidationErrors(err)
		return resp.ValidationErrorResponse(c, "failed to validate request", validationErrors)
	}

	otp := h.service.GenerateOTP()

	err = h.service.CreateNewCustomer(context.Background(), request.customer, otp)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "error while inserting user", err)
	}
	switch strings.ToLower(request.OTPDeliveryMethod) {
	case "sms":
		if _, err := h.service.SendSMS(request.customer.PhoneNumber, otp); err != nil {
			return resp.ErrorResponse(c, http.StatusInternalServerError, "failed to send SMS OTP", err)
		}
	case "email":
		if err := h.service.SendEmail(request.customer.Email, otp); err != nil {
			return resp.ErrorResponse(c, http.StatusInternalServerError, "failed to send email OTP", err)
		}
	}

	request.customer.Password = ""
	return resp.SuccessResponse(c, http.StatusOK, "Check your "+request.OTPDeliveryMethod+" for verification code",
		request.customer.Email)
}

func (h *CustomerHandler) LoginUser(c echo.Context) error {
	type CustomerLogin struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=8"`
	}

	var customer CustomerLogin
	if err := c.Bind(&customer); err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}

	if err := h.service.Validate(customer); err != nil {
		validationErrors := h.service.GetValidationErrors(err)
		return resp.ValidationErrorResponse(c, "failed to validate request", validationErrors)
	}
	
	token, err := h.service.LoginCustomer(context.Background(), customer.Email, customer.Password)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "", err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"access_token": token,
	})
}

func (h *CustomerHandler) VerifyOTP(c echo.Context) error {

}
