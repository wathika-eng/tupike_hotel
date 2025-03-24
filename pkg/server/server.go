package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"tupike_hotel/pkg/config"
	"tupike_hotel/pkg/database"
	"tupike_hotel/pkg/migration"
	"tupike_hotel/pkg/routes"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	port    string
	db      database.DBService
	client  *redis.Client
	handler http.Handler
}

func NewServer() *http.Server {
	// Initialize Database
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Initialize Redis Client
	client := redis.NewClient(&redis.Options{
		Addr:     config.Envs.RedisUrl,
		Password: "", // No password set
		DB:       0,  // Default DB
		Protocol: 2,  // Connection protocol
	})

	// Check Redis Connection
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Printf("connected to redis :%v", client)
	// Initialize Server
	server := &Server{
		port:    config.Envs.ServerPort,
		db:      db,
		client:  client,
		handler: routes.SetupRoutes(db, client),
	}

	// Run Migrations
	if err := migration.Migrate(db.GetDB()); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Printf("Server running on http://localhost:%v\n", server.port)

	// Return HTTP Server
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", server.port),
		Handler:      server.handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
