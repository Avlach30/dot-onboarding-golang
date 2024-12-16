package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gorm.io/gorm"
)

type RoleSeeder struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

// Handle implements Seeder.
func (roleSeeder *RoleSeeder) Handle(db *gorm.DB) error {
	filePath := "seeder/files/roles.json"

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
	var roles []RoleSeeder
	if err := json.Unmarshal(bytes, &roles); err != nil {
		return fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	// Insert each role into the database
	db.Exec(`DELETE FROM role_entities`)
	for _, role := range roles {
		log.Println(role)
		db.Exec(`INSERT INTO role_entities (name, key) VALUES (?, ?)`, role.Name, role.Key)
	}

	return nil
}

func NewRoleSeeder() Seeder {
	return &RoleSeeder{}
}
