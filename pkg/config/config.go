// provides secret config keys for our API
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	ServerPort       string
	DbType           string
	DbHost           string
	DbPort           string
	DbUser           string
	DbPassword       string
	DbName           string
	ConnectionString string
	SecretKey        string
	RefreshKey       string
	RESEND_API_KEY   string
	// UPTRACE_DSN       string
	GinMode            string
	RedisUrl           string
	DbMaxIdleConns     string
	DbMaxOpenConns     string
	DbMaxLifetimeConns string
	AtApiKey           string
	AtUserName         string
	AtShortCode        string
	AtEnvironment      string
}

var Envs = initConfigs()

func initConfigs() config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Warning: No .env file found or failed to load. Using default environment variables.")
	}
	return config{
		ServerPort:       getEnv("SERVER_PORT", "8000"),
		DbType:           getEnv("DB_TYPE", "postgresql"),
		DbHost:           getEnv("DB_HOST", "localhost"),
		DbPort:           getEnv("DB_PORT", "5432"),
		DbUser:           getEnv("DB_USER", "postgres"),
		DbPassword:       getEnv("DB_PASSWORD", "postgres"),
		DbName:           getEnv("DB_NAME", "todoApp"),
		ConnectionString: getEnv("CONNECTION_STRING", ""),
		SecretKey:        getEnv("SECRET_KEY", "https://acte.ltd/utils/randomkeygen"),
		RefreshKey:       getEnv("REFRESH_KEY", "https://randomkeygen.com/"),
		RESEND_API_KEY:   getEnv("RESEND_API_KEY", ""),
		// UPTRACE_DSN:       getEnv("UPTRACE_DSN", ""),
		GinMode:            getEnv("GIN_MODE", "release"),
		RedisUrl:           getEnv("REDIS_URL", "localhost:6379"),
		DbMaxIdleConns:     getEnv("DB_MAX_IDLE_CONNS", "10"),
		DbMaxOpenConns:     getEnv("DB_MAX_OPEN_CONNS", "10"),
		DbMaxLifetimeConns: getEnv("DB_MAX_LIFETIME_CONNS", "10"),
		AtApiKey:           getEnv("AtApiKey", ""),
		AtUserName:         getEnv("AtUserName", "sandbox"),
		AtShortCode:        getEnv("AtShortCode", ""),
		AtEnvironment:      getEnv("AtEnvironment", "sandbox"),
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
