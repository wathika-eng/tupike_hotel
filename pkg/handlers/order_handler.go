package handlers

import (
	"errors"
	"net/http"
	resp "tupike_hotel/pkg/response"
	"tupike_hotel/pkg/types"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CustomerID(c echo.Context) (uuid.UUID, error) {
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errors.New("unauthorized")
	}
	customerID, _ := claims["sub"].(uuid.UUID)
	return customerID, nil
}

func (h *Handler) OrderFood(c echo.Context) error {
	customerID, err := CustomerID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	var order types.Order
	order.CustomerID = customerID
	if err := c.Bind(&order); err != nil {
		return resp.ErrorResponse(c, 400, "", err)
	}

	// if err := h.service.Validate(&order); err != nil {
	// 	return resp.ErrorResponse(c, 400, "", err)
	// }
	err = h.service.PlaceOrder(c.Request().Context(), &order)
	if err != nil {
		return resp.ErrorResponse(c, 400, "", err)
	}
	return resp.SuccessResponse(c, http.StatusCreated, "order placed succesfully", order)
}
