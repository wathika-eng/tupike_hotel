package resp

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

// Response represents a standardized API response.
type Response struct {
	Code    int    `json:"code"`              // HTTP status code
	Message string `json:"message,omitempty"` // Optional message for success/error
	Data    any    `json:"data,omitempty"`    // Response data for success
	Error   string `json:"error,omitempty"`   // Error message for failure
}

// SuccessResponse sends a standardized success response.
func SuccessResponse(c echo.Context, code int, message string, data any) error {
	response := Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return c.JSON(code, response)
}

// ErrorResponse sends a standardized error response.
func ErrorResponse(c echo.Context, code int, message string, err error) error {
	// Log the error for debugging
	log.Errorf("API error: %v", err)

	response := Response{
		Code:    code,
		Message: message,
		Error:   err.Error(),
	}
	return c.JSON(code, response)
}

// ValidationErrorResponse sends a standardized response for validation errors.
func ValidationErrorResponse(c echo.Context, message string, validationErrors map[string]string) error {
	response := Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    validationErrors, // Include validation errors in the response
	}
	return c.JSON(http.StatusBadRequest, response)
}
