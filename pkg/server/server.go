package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"tupike_hotel/pkg/config"
	"tupike_hotel/pkg/database"
)

type Server struct {
	port string
	db   database.DBService
}

// type ServerInterface interface {
// 	healthChecker()
// }

func NewServer() *http.Server {
	NewServer := &Server{
		port: config.Envs.SERVER_PORT,
		db: func() database.DBService {
			db, err := database.NewDatabase()
			if err != nil {
				log.Fatalf("error connecting to the database: %v", err.Error())
			}
			return db
		}(),
	}
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", NewServer.port),
		Handler:      NewServer.SetupRoutes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
