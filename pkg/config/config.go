// provides secret config keys for our API
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	SERVER_PORT       string
	DB_TYPE           string
	DB_HOST           string
	DB_PORT           string
	DB_USER           string
	DB_PASSWORD       string
	DB_NAME           string
	CONNECTION_STRING string
	SECRET_KEY        string
	REFRESH_KEY       string
	// RESEND_API_KEY    string
	// UPTRACE_DSN       string
	GIN_MODE              string
	REDIS_URL             string
	DB_MAX_IDLE_CONNS     string
	DB_MAX_OPEN_CONNS     string
	DB_MAX_LIFETIME_CONNS string
}

var Envs = initConfigs()

func initConfigs() config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found or failed to load. Using default environment variables.")
	}
	return config{
		SERVER_PORT:       getEnv("SERVER_PORT", "8000"),
		DB_TYPE:           getEnv("DB_TYPE", "postgresql"),
		DB_HOST:           getEnv("DB_HOST", "localhost"),
		DB_PORT:           getEnv("DB_PORT", "5432"),
		DB_USER:           getEnv("DB_USER", "postgres"),
		DB_PASSWORD:       getEnv("DB_PASSWORD", "postgres"),
		DB_NAME:           getEnv("DB_NAME", "todoApp"),
		CONNECTION_STRING: getEnv("CONNECTION_STRING", ""),
		SECRET_KEY:        getEnv("SECRET_KEY", "https://acte.ltd/utils/randomkeygen"),
		REFRESH_KEY:       getEnv("REFRESH_KEY", "https://randomkeygen.com/"),
		// RESEND_API_KEY:    getEnv("RESEND_API_KEY", ""),
		// UPTRACE_DSN:       getEnv("UPTRACE_DSN", ""),
		GIN_MODE:              getEnv("GIN_MODE", "release"),
		REDIS_URL:             getEnv("REDIS_URL", "localhost:6379"),
		DB_MAX_IDLE_CONNS:     getEnv("DB_MAX_IDLE_CONNS", "10"),
		DB_MAX_OPEN_CONNS:     getEnv("DB_MAX_OPEN_CONNS", "10"),
		DB_MAX_LIFETIME_CONNS: getEnv("DB_MAX_LIFETIME_CONNS", "10"),
	}
}

// getEnv retrieves an environment variable or returns the fallback string
func getEnv(checkKey, fallback string) string {
	key, ok := os.LookupEnv(checkKey)
	if !ok {
		log.Printf("%v key not found, using default value", checkKey)
		return fallback
	}
	return key
}
