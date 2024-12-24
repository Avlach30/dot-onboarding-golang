package migration

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type SchemaMigration struct {
	Version int64
}

func Create(db *gorm.DB, fileName string) {
	if fileName != "" {
		now := time.Now().Format("20060102150405")
		upFileName := fmt.Sprintf("migration/files/%s_%s.up.sql", now, fileName)
		downFileName := fmt.Sprintf("migration/files/%s_%s.down.sql", now, fileName)

		if err := os.WriteFile(upFileName, []byte(""), 0644); err != nil {
			log.Fatalf("failed to create up migration file : %v", err)
		}

		if err := os.WriteFile(downFileName, []byte(""), 0644); err != nil {
			log.Fatalf("failed to create down migration file : %v", err)
		}

		fmt.Printf("Created migration files:\n%s\n%s\n", upFileName, downFileName)
	}
}

func Run(db *gorm.DB, exec string) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"") // Create UUID extension

	// Extract raw SQL DB from GORM
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get raw DB from GORM : %v", err)
	}

	// Initialize migrate with the PostgreSQL driver
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not create postgres driver : %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration/files", // Path to migration files
		"postgres",               // Database name
		driver,
	)
	if err != nil {
		log.Fatalf("could not create migrate instance : %v", err)
	}

	// Check if migration is dirty and force it if needed
	version, dirty, _ := m.Version()
	if dirty {
		fmt.Printf("Migration is dirty. Forcing version %d\n", version)
		if err := m.Force(int(version)); err != nil {
			log.Fatalf("failed to force migration version : %v", err)
		}
	}

	// Execute migration based on user input
	switch exec {
	case "down":
		err = m.Steps(-1)
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run down migration : %v", err)
		}
	case "fresh":
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run down all migration : %v", err)
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run up migration : %v", err)
		}
	default:
		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			log.Fatalf("failed to run up migration : %v", err)
		}
	}

	fmt.Println("Migration applied successfully!")
}
