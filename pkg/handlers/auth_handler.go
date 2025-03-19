package handlers

import (
	"context"
	"net/http"
	resp "tupike_hotel/pkg/response"
	"tupike_hotel/pkg/types"

	"github.com/labstack/echo/v4"
)

func (h *CustomerHandler) CreateUser(c echo.Context) error {
	var customer types.Customer
	err := c.Bind(&customer)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}
	err = h.service.Validate(customer)
	if err != nil {
		validationErrors := h.service.GetValidationErrors(err)
		return resp.ValidationErrorResponse(c, "failed to validate request", validationErrors)
	}
	err = h.service.CreateNewCustomer(context.Background(), &customer)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "error while inserting user", err)
	}
	customer.Password = ""
	return resp.SuccessResponse(c, http.StatusOK, "user created successfully", customer.Email)
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
	return resp.SuccessResponse(c, http.StatusOK, "user signed in", token)
}
