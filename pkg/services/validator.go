package services

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func (s *Service) Validate(i any) error {
	if err := s.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func (s *Service) GetValidationErrors(err error) map[string]string {
	validationErrors := make(map[string]string)

	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			validationErrors[fe.Field()] = fmt.Sprintf(
				"The field '%s' failed validation with the rule '%s'",
				fe.Field(),
				fe.Tag(),
			)
		}
	} else {
		validationErrors["error"] = err.Error()
	}

	return validationErrors
}
