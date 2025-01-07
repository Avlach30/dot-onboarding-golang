package dbconn

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseCredentials struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
	TimeZome string
}

func InitDb(databaseCredentials *DatabaseCredentials) (*gorm.DB, error) {

	connectionString := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s"
	dsn := fmt.Sprintf(
		connectionString,
		databaseCredentials.Host,
		databaseCredentials.Username,
		databaseCredentials.Password,
		databaseCredentials.Name,
		databaseCredentials.Port,
		databaseCredentials.TimeZome,
	)

	db, errors := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if errors != nil {
		panic(errors.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println("Failed to get raw DB from GORM", err)
		panic(*exception.ServerErrorException(err))
	}

	// Set connection pool configuration
	maxIdleConn, err := strconv.Atoi(config.DBMaxIdleConn)
	if err != nil {
		log.Fatalf("Failed to convert DBMaxIdleConn to int : %v", err)
	}
	maxOpenConn, err := strconv.Atoi(config.DBMaxOpenConn)
	if err != nil {
		log.Fatalf("Failed to convert DBMaxOpenConn to int : %v", err)
	}
	connMaxIdleTime, err := strconv.Atoi(config.DBConnMaxIdleTime)
	if err != nil {
		log.Fatalf("Failed to convert DBConnMaxIdleTime to int : %v", err)
	}
	connMaxLifetime, err := strconv.Atoi(config.DBConnMaxLifetime)
	if err != nil {
		log.Fatalf("Failed to convert DBConnMaxLifetime to int : %v", err)
	}

	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxOpenConn)
	sqlDB.SetConnMaxIdleTime(time.Duration(connMaxIdleTime) * time.Minute)
	sqlDB.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)

	log.Println("DBMaxIdleConn : ", maxIdleConn)
	log.Println("DBMaxOpenConn : ", maxOpenConn)
	log.Println("DBConnMaxIdleTime : ", connMaxIdleTime)
	log.Println("DBConnMaxLifetime : ", connMaxLifetime)

	return db, errors
}
