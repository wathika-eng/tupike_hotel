package handlers

import (
	"net/http"
	resp "tupike_hotel/pkg/response"
	"tupike_hotel/pkg/types"

	"github.com/labstack/echo/v4"
)

// adds food to the database, only admin can perform this
func (h *Handler) AddFood(c echo.Context) error {
	var food types.FoodItem
	if err := c.Bind(&food); err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "", err)
	}
	if err := h.service.Validate(food); err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "", err)
	}

	if err := h.service.AddFood(c.Request().Context(), &food); err != nil {
		return resp.ErrorResponse(c, http.StatusBadGateway, "", err)
	}
	return resp.SuccessResponse(c, http.StatusCreated, "created successfully", food)
}

// GetFood fetched food from the data and returns a json response
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
