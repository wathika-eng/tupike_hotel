package handlers

import (
	"net/http"
	"tupike_hotel/pkg/repository"
	resp "tupike_hotel/pkg/response"
	"tupike_hotel/pkg/services"
	"tupike_hotel/pkg/types"

	"github.com/labstack/echo/v4"
)

type CustomerHandler struct {
	repo    repository.RepoInterface
	service services.ServiceInterface
}

type CustomerHandlerInterface interface {
	CreateUser(e echo.Context) error
}

func NewCustomerHandler(repo repository.RepoInterface,
	service services.ServiceInterface) CustomerHandlerInterface {
	return &CustomerHandler{
		repo:    repo,
		service: service,
	}
}

func (h *CustomerHandler) CreateUser(c echo.Context) error {
	var customer types.Customer
	err := c.Bind(&customer)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusBadRequest, "invalid request body", err)
	}
	err = h.service.Validate(customer)
	if err != nil {
		validationErrors := h.service.GetValidationErrors(err)
		return resp.ValidationErrorResponse(c, "failed to validate request", validationErrors)
	}
	return resp.SuccessResponse(c, http.StatusOK, "user created successfully", customer)
}

func (h *CustomerHandler) HealthChecker(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"Status":  http.StatusOK,
		"Results": "",
	})
}
