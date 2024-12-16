package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/gorm"
)

type PermissionSeeder struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Handle implements Seeder.
func (permissionSeeder *PermissionSeeder) Handle(db *gorm.DB) error {
	filePath := "seeder/files/permissions.json"

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open JSON file: %v", err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read JSON file: %v", err)
	}

	// Parse JSON data
	var permissions []PermissionSeeder
	if err := json.Unmarshal(bytes, &permissions); err != nil {
		return fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	// Insert each permission into the database
	for _, permission := range permissions {
		log.Println(permission)
		db.Exec(`DELETE FROM permission_entities`)
		db.Exec(`INSERT INTO permission_entities (name, key) VALUES (?, ?)`, permission.Name, permission.Key)
	}

	return nil
}

func NewPermissionSeeder() Seeder {
	return &PermissionSeeder{}
}
