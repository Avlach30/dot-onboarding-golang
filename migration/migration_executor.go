package migration

import (
	"log"

	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, entities []interface{}) {
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	errors := db.AutoMigrate(entities...)
	if errors != nil {
		panic(errors.Error())
	}

	log.Println("auto migration complete")
}
