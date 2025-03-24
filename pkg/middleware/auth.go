package middleware

import (
	"log"
	"net/http"
	"strings"
	"tupike_hotel/pkg/services"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				log.Println("Missing Authorization header")
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
			}

			// Expect "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				log.Println("Invalid Authorization header format")
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization header format")
			}

			tokenString := parts[1]
			claims, err := services.VerifyToken(tokenString)
			if err != nil {
				log.Println("JWT verification failed:", err)
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
			}

			// Store claims in context
			c.Set("claims", claims)
			log.Println("Token verified, claims:", claims)

			return next(c)
		}
	}
}
