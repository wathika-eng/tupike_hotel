package server

import (
	"net/http"
	logger "tupike_hotel/pkg/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// type response struct {
// 	Status  int               `json:"status"`
// 	Results map[string]string `json:"results"`
// }

func (s *Server) SetupRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	e.Use(logger.LoggerMiddleware)
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "DELETE", "POST", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	e.GET("/", s.healthChecker)

	return e
}

func (s *Server) healthChecker(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"Status":  http.StatusOK,
		"Results": s.db.Health(),
	})
}
