package migration

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
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
			fmt.Println("Failed to create up migration file", err)
			panic(*exception.ServerErrorException(err))
		}

		if err := os.WriteFile(downFileName, []byte(""), 0644); err != nil {
			fmt.Println("Failed to create down migration file", err)
			panic(*exception.ServerErrorException(err))
		}

		fmt.Printf("Created migration files:\n%s\n%s\n", upFileName, downFileName)
	}
}

func Run(db *gorm.DB, exec string) {
	// Create UUID extension
	err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		fmt.Println("Failed to create UUID extension", err)
		panic(*exception.ServerErrorException(err))
	}

	// Extract raw SQL DB from GORM
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("Failed to get raw DB from GORM", err)
		panic(*exception.ServerErrorException(err))
	}

	// Initialize migrator
	migrator, err := initializeMigrator(sqlDB)
	if err != nil {
		fmt.Println("Failed to initialize migrator", err)
		panic(*exception.ServerErrorException(err))
	}

	// Handle dirty migrations
	handleDirtyMigration(migrator)
	// Execute migration based on user input
	err = executeMigration(migrator, exec)
	if err != nil {
		fmt.Println("Migration failed", err)
		panic(*exception.ServerErrorException(err))
	}
	fmt.Println("Migration applied successfully!")

}

func initializeMigrator(sqlDB *sql.DB) (*migrate.Migrate, error) {
	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		fmt.Println("Failed to create postgres driver", err)
		return nil, fmt.Errorf("could not create postgres driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file://migration/files", "postgres", driver)
	if err != nil {
		fmt.Println("Failed to create migrate instance", err)
		return nil, fmt.Errorf("could not create migrate instance: %v", err)
	}

	return m, nil
}

func handleDirtyMigration(m *migrate.Migrate) {
	version, dirty, _ := m.Version()
	if dirty {
		fmt.Printf("Migration is dirty. Forcing version %d\n", version)
		if err := m.Force(int(version)); err != nil {
			fmt.Println("Failed to force migration version", err)
			panic(*exception.ServerErrorException(err))
		}
	}

}

func executeMigration(m *migrate.Migrate, exec string) error {
	var err error

	switch exec {
	case "down":
		err = m.Steps(-1)
		if err != nil && err != migrate.ErrNoChange {
			fmt.Println("Failed to run down migration", err)
			panic(*exception.ServerErrorException(err))
		}
	case "fresh":
		err = m.Down()
		if err != nil && err != migrate.ErrNoChange {
			fmt.Println("Failed to run down all", err)
			panic(*exception.ServerErrorException(err))
		}

		err = m.Up()
		if err != nil && err != migrate.ErrNoChange {
			fmt.Println("Failed to run up migration", err)
			panic(*exception.ServerErrorException(err))
		}
	default:
		err = m.Up()
	}

	if err == migrate.ErrNoChange {
		fmt.Println("No migration changes")
		err = nil
	}

	return err
}
