package handlers

import (
	"context"
	"net/http"
	resp "tupike_hotel/pkg/response"
	"tupike_hotel/pkg/types"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateUser(c echo.Context) error {

	var customer *types.Customer

	err := c.Bind(&customer)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}
	err = h.service.Validate(customer)
	if err != nil {
		validationErrors := h.service.GetValidationErrors(err)
		return resp.ValidationErrorResponse(c, "failed to validate request", validationErrors)
	}

	otp := h.service.GenerateOTP()
	customer.OTP = otp
	//send OTP to both email and phone number
	if _, err := h.service.SendOTP(c.Request().Context(), customer.PhoneNumber, customer.Email, otp); err != nil {
		return resp.ErrorResponse(c, http.StatusInternalServerError, "failed to send OTP", err)
	}

	err = h.service.CreateNewCustomer(context.Background(), customer)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "error while inserting user", err)
	}

	customer.Password = ""
	return resp.SuccessResponse(c, http.StatusOK, "Check your email or phone number for verification code",
		customer.Email)
}

func (h *Handler) LoginUser(c echo.Context) error {
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

func (h *Handler) VerifyOTP(c echo.Context) error {
	var OTP struct {
		Email string `json:"email"`
		Code  string `json:"code" validate:"required,max=4"`
	}
	if err := c.Bind(&OTP); err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}
	if err := h.service.Validate(OTP); err != nil {
		validationErrors := h.service.GetValidationErrors(err)
		return resp.ValidationErrorResponse(c, "failed to validate request", validationErrors)
	}
	err := h.service.CheckOTP(c.Request().Context(), OTP.Email, OTP.Code)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "", err)
	}
	return resp.SuccessResponse(c, http.StatusOK, "otp verified successfully", nil)
}
