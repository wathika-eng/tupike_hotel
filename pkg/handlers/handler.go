package handlers

import (
	"net/http"
	"tupike_hotel/pkg/repository"
	"tupike_hotel/pkg/services"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	repo    *repository.Repository
	service *services.Service
}

func NewHandler(repo *repository.Repository,
	service *services.Service) *Handler {
	return &Handler{
		repo:    repo,
		service: service,
	}
}

// type AuthHandler interface {
// 	CreateUser(c echo.Context) error
// 	VerifyOTP(c echo.Context) error
// 	LoginUser(c echo.Context) error
// }

// type UserHandler interface {
// 	Profile(c echo.Context) error
// }

// type FoodHandler interface {
// 	AddFood(c echo.Context) error
// 	GetFood(c echo.Context) error
// }

// type OrderHandler interface {
// 	OrderFood(c echo.Context) error
// }

func (h *Handler) HealthChecker(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Status":  http.StatusOK,
		"Results": "",
	})
}
