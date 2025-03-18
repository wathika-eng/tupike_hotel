package server

import (
	"net/http"
	logger "tupike_hotel/pkg/middleware"
	"tupike_hotel/pkg/repository"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) SetupRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	e.Use(logger.LoggerMiddleware)
	repo := repository.NewRepository()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "DELETE", "POST", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	e.GET("/", s.healthChecker)
	api := e.Group("/api")
	api.POST("/user")
	return e
}

func (s *Server) healthChecker(c echo.Context) error {

	return c.JSON(http.StatusOK, echo.Map{
		"Status":  http.StatusOK,
		"Results": s.db.Health(),
	})
}
