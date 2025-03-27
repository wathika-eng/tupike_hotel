package handlers

import (
	"net/http"
	resp "tupike_hotel/pkg/response"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Profile(c echo.Context) error {
	claims, ok := c.Get("claims").(jwt.MapClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusForbidden, "Unauthorized access")
	}
	email, ok := claims["sub"].(string)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Email not found in token")
	}
	user, err := h.repo.CustomerRepo.LookUpCustomer(c.Request().Context(), email)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusInternalServerError,
			"user not found", err)
	}
	user.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}
