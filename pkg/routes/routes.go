package routes

import (
	"net/http"
	"tupike_hotel/pkg/database"
	"tupike_hotel/pkg/handlers"
	logger "tupike_hotel/pkg/middleware"
	"tupike_hotel/pkg/repository"
	"tupike_hotel/pkg/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetupRoutes(db database.DBService) http.Handler {
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
	repo := repository.NewRepository(db.GetDB())
	services := services.NewService(repo, validator.New())
	handler := handlers.NewCustomerHandler(repo, services)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"Status":  http.StatusOK,
			"Results": db.Health(),
		})
	})

	api := e.Group("/api")
	{
		api.POST("/signup", handler.CreateUser)
		api.POST("/login", handler.LoginUser)
	}
	return e
}
