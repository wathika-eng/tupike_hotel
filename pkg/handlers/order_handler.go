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

// logic to extract user ID from token
func CustomerID(c echo.Context) (string, error) {
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return "", errors.New("unauthorized")
	}
	// (string)
	sub, ok := claims["sub"].(string)
	if !ok {
		return "", errors.New("invalid subject type")
	}

	customerID, err := uuid.Parse(sub)
	if err != nil {
		return "", errors.New("invalid uuid in sub")
	}

	return customerID.String(), nil
}

func (h *Handler) OrderFood(c echo.Context) error {
	customerID, err := CustomerID(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "unauthorized")
	}
	id, _ := uuid.Parse(customerID)

	var order types.Order
	order.CustomerID = id

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
