package dbconn

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	db, errors := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if errors != nil {
		panic(errors.Error())
	}

	return db, errors
}
