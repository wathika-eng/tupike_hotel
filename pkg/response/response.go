package response

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code  int    `json:"code"`
	Data  any    `json:"data"`
	Error string `json:"error,omitempty"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func JSONResponse(c echo.Context, code int, data any, err error) error {
	response := Response{Code: code}
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Data = data
	}
	return c.JSON(code, response)
}
