package migration

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, exec string) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS 'uuid-ossp'")

	// Extract raw SQL DB from GORM
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get raw DB from GORM: %v", err)
	}

	// Initialize migrate with the PostgreSQL driver
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration/files", // Path to migration files
		"postgres",               // Database name
		driver,
	)
	if err != nil {
		log.Fatalf("could not create migrate instance: %v", err)
	}

	// Execute migration based on user input
	switch exec {
	case "down":
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run down migration: %v", err)
		}
	default:
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run up migration: %v", err)
		}
	}

	fmt.Println("Migration applied successfully!")
}
