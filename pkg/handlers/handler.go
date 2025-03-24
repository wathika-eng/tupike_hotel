package handlers

import (
	"net/http"
	"tupike_hotel/pkg/repository"
	"tupike_hotel/pkg/services"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	repo    repository.RepoInterface
	service services.ServiceInterface
}

type CustomerHandlerInterface interface {
	CreateUser(e echo.Context) error
	LoginUser(e echo.Context) error
	AddFood(c echo.Context) error
	OrderFood(e echo.Context) error
	Profile(c echo.Context) error
}

func NewCustomerHandler(repo repository.RepoInterface,
	service services.ServiceInterface) CustomerHandlerInterface {
	return &CustomerHandler{
		repo:    repo,
		service: service,
	}
}

func (h *CustomerHandler) HealthChecker(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Status":  http.StatusOK,
		"Results": "",
	})
}
