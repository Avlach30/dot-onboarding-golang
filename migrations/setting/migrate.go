package main

import (
	"fmt"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/common/enum"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/dbconn"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	pathToMigrationsFile := "file://migrations"

	db, err := dbconn.GetDb(enum.MYSQL)

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		panic(fmt.Sprintf("error in fetching driver from sql DB: %v", err))
	}

	m, err := migrate.NewWithDatabaseInstance(pathToMigrationsFile, config.Database, driver)
	if err != nil {
		panic(fmt.Sprintf("failed to create migration instance: %v", err))
	}

	err = m.Up()
	if err != nil {
		if err.Error() != "no change" {
			panic(fmt.Sprintf("failed to up the migration: %v", err))
		}
		fmt.Println("No change from previous migration")
	}

	fmt.Println("Migration process complete")
}
