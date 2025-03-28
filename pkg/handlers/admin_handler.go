package handlers

import (
	"net/http"
	resp "tupike_hotel/pkg/response"

	"github.com/labstack/echo/v4"
)

// returns a users profile
// extract id from the token then search on the DB
// implement cache later
func (h *Handler) Profile(c echo.Context) error {
	customerID, err := CustomerID(c)
	user, err := h.repo.CustomerRepo.LookUpCustomer(c.Request().Context(), customerID)
	if err != nil {
		return resp.ErrorResponse(c, http.StatusInternalServerError,
			"user not found", err)
	}
	user.Password = ""
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}
