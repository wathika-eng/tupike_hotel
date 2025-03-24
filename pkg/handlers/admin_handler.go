package handlers

import (
	"log"
	"net/http"
	resp "tupike_hotel/pkg/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// can perform crud operations on food items
// only admins
func (h *CustomerHandler) AddFood(c echo.Context) error {
	return nil
}

func (h *CustomerHandler) Profile(c echo.Context) error {
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusForbidden, "Unauthorized access")
	}

	log.Println("Claims in handler:", claims)

	return resp.SuccessResponse(c, 200, "testing", nil)
}
