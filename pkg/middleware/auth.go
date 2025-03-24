package middleware

import (
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// todo
func AuthMiddleware() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Retrieve token from context
		token, ok := c.Get("user").(*jwt.Token)
		if !ok || token == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "JWT token missing or invalid")
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return echo.NewHTTPError(http.StatusUnauthorized, "Failed to parse JWT claims")
		}

		// Store claims in context for further use
		c.Set("claims", claims)
		log.Printf("%v\n", claims)
		// Allow the request to proceed
		return nil
	}
}
