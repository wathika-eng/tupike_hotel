package resp

import (
	"github.com/labstack/echo/v4"
)

type Response struct {
	Code  int    `json:"code"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// JSONResponse standardizes API responses.
func JSONResponse(c echo.Context, code int, data any, err error) error {
	response := Response{Code: code}
	if err != nil {
		response.Error = err.Error()
	} else {
		response.Data = data
	}
	return c.JSON(code, response)
}
