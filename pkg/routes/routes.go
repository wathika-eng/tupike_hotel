package routes

import (
	"net/http"
	"tupike_hotel/pkg/database"
	"tupike_hotel/pkg/handlers"
	custom "tupike_hotel/pkg/middleware"
	"tupike_hotel/pkg/repository"
	"tupike_hotel/pkg/services"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/redis/go-redis/v9"
)

func SetupRoutes(db database.DBService, client *redis.Client) http.Handler {
	e := echo.New()
	e.Use(middleware.Recover())
	// e.Use(middleware.Logger())
	//e.Use(echojwt.JWT([]byte(config.Envs.SecretKey)))
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "DELETE", "POST", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	repo := repository.NewRepository(repository.NewDatabaseManager(db.GetDB(), client))
	services := services.NewService(repo.CustomerRepo, repo.FoodRepo, repo.OrderRepo, validator.New())
	handler := handlers.NewHandler(repo, services)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status":  http.StatusOK,
			"results": db.Health(),
		})
	})

	api := e.Group("/api")
	{
		api.GET("/food", handler.GetFood)
	}

	apiAuth := e.Group("/api/auth")
	{
		apiAuth.POST("/signup", handler.CreateUser)
		apiAuth.POST("/verify-otp", handler.VerifyOTP)
		//api.POST("/generate-otp")
		apiAuth.POST("/login", handler.LoginUser)
	}

	r := e.Group("/api/protected")
	r.Use(custom.AuthMiddleware())
	{
		r.GET("/profile", handler.Profile)
		r.POST("/add-food", handler.AddFood)
		r.POST("/place-order", handler.OrderFood)
	}

	return e
}
