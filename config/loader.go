package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file : %v", err)
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
