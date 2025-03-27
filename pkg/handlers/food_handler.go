package handlers

import (
	"net/http"
	resp "tupike_hotel/pkg/response"

	"github.com/labstack/echo/v4"
)

func (h *Handler) GetFood(c echo.Context) error {
	food, err := h.service.FetchFood(c.Request().Context())
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadGateway, "error", err)
	}
	if len(food) <= 0 {
		return c.JSON(200, echo.Map{
			"message": "no food fetched",
			"food":    food,
		})
	}
	return c.JSON(200, echo.Map{
		"food": food,
	})
}
