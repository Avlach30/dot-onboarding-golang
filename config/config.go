package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func InitConfig() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	return nil
}

func GetAppModeEnv() string {
	appMode := os.Getenv("APP_MODE")

	if appMode == "" {
		appMode = "development"
	}

	return appMode
}

type DBConfig struct {
	Host     string
	Username string
	Password string
	Db       string
}

func GetMySqlEnv() DBConfig {
	dbHost := os.Getenv("MYSQL_DB_HOST")
	dbUser := os.Getenv("MYSQL_DB_USER")
	dbPass := os.Getenv("MYSQL_DB_PASS")
	dbName := os.Getenv("MYSQL_DB_NAME")

	if dbHost == "" || dbUser == "" || dbPass == "" || dbName == "" {
		log.Fatal("GetMySqlEnv.EnvNotDefined")
	}

	return DBConfig{
		Host:     dbHost,
		Username: dbUser,
		Password: dbPass,
		Db:       dbName,
	}
}
