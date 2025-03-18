package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"tupike_hotel/pkg/config"
	"tupike_hotel/pkg/database"
	"tupike_hotel/pkg/routes"
)

type Server struct {
	port    string
	db      database.DBService
	handler http.Handler
}

func NewServer() *http.Server {
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err.Error())
	}
	NewServer := &Server{
		port:    config.Envs.ServerPort,
		db:      db,
		handler: routes.SetupRoutes(db),
	}
	//migrations.Migrate(NewServer.db.GetDB())
	//migrations.Drop(NewServer.db.GetDB())
	defer log.Printf("serving on http://localhost:%v\n", NewServer.port)
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", NewServer.port),
		Handler:      NewServer.handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
