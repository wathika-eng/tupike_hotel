package handlers

import (
	"errors"
	"net/http"
	"strings"
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

// search for food using UUID or food name, return all it's data
func (h *Handler) SearchFood(c echo.Context) error {
	type request struct {
		Name string `json:"name" validete:"required"`
	}
	var req request
	name := c.QueryParam("food")
	if strings.TrimSpace(name) == "" {
		if err := c.Bind(&req); err != nil {
			return resp.ErrorResponse(c, http.StatusBadRequest, "", errors.New("request body cannot be empty"))
		}
		name = req.Name
	}
	if strings.TrimSpace(name) == "" {
		return resp.ErrorResponse(c, http.StatusBadRequest, "", errors.New("request body cannot be empty"))
	}
	foodData, err := h.service.CheckFood(c.Request().Context(), name)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "", err)
	}
	return resp.SuccessResponse(c, http.StatusOK, "food found", foodData)
}
