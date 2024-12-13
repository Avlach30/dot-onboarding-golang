package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	return nil
}

func Get(key, defaultValue string) string {
	LoadConfig()

	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetRequired(key string) string {
	LoadConfig()

	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}

var (
	// General
	AppMode = Get("APP_MODE", "PROD")
	GinMode = Get("GIN_MODE", "release")

	AppPort = Get("APP_PORT", "8080")

	SentryDSN         = Get("SENTRY_DSN", "")
	SentrySampleTrace = Get("SENTRY_SAMPLE_TRACE", "0.1")

	Secret           = GetRequired("JWT_SECRET")
	JwtExpiredInDays = GetRequired("JWT_EXPIRED_IN_DAYS")
	ServiceAuthToken = GetRequired("SERVICE_AUTH_TOKEN")

	// DATABASE
	DBHost     = GetRequired("DB_HOST")
	DBPort     = GetRequired("DB_PORT")
	DBUsername = GetRequired("DB_USER")
	DBPassword = GetRequired("DB_PASS")
	DBName     = GetRequired("DB_NAME")
	DBTimeZone = GetRequired("DB_TIMEZONE")

	// REDIS
	RedisHost     = Get("REDIS_HOST", "")
	RedisPassword = Get("REDIS_Password", "")
	RedisPort     = Get("REDIS_PORT", "")
)
