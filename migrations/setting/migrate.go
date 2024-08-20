package main

import (
	"fmt"
	"github.com/codespace-id/codespace-x/config"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/codespace-id/codespace-x/pkg/dbconn"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	if config.AppMode != "MIGRATE" {
		log.Fatalf("Please recheck your DB & adjust appMode to MIGRATE")
	}

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
