package handlers

import (
	"errors"
	"tupike_hotel/pkg/repository"
	resp "tupike_hotel/pkg/response"
	"tupike_hotel/pkg/services"
	"tupike_hotel/pkg/types"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	repo    repository.RepoInterface
	service services.CustomerServiceInterface
}

type CustomerHandlerInterface interface {
	CreateUser(e echo.Context) error
}

func NewCustomerHandler(repo repository.RepoInterface,
	service services.CustomerServiceInterface) CustomerHandlerInterface {
	return &CustomerHandler{
		repo:    repo,
		service: service,
	}
}

func (h *CustomerHandler) CreateUser(c echo.Context) error {
	var customer types.Customer
	err := c.Bind(customer)
	if err != nil {
		return resp.JSONResponse(c, 400, nil, errors.New(err.Error()))
	}
	return resp.JSONResponse(c, 201, "okay", nil)
}
