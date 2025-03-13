package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type response struct {
// 	Status  int               `json:"status"`
// 	Results map[string]string `json:"results"`
// }

func (s *Server) SetupRoutes() http.Handler {
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"GET", "DELETE", "POST", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	r.GET("/", s.healthChecker)

	return r
}

func (s *Server) healthChecker(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"Status":  http.StatusOK,
		"Results": s.db.Health(),
	})
}
