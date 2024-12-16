package seeder

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserSeeder struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Handle implements Seeder.
func (userSeeder *UserSeeder) Handle(db *gorm.DB) error {
	filePath := "seeder/files/users.json"

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
	var users []UserSeeder
	if err := json.Unmarshal(bytes, &users); err != nil {
		return fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	// Insert each user into the database
	for _, user := range users {
		log.Println(user)
		passwordByte, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
		password := string(passwordByte)
		db.Exec(`DELETE FROM user_entities`)
		db.Exec(`INSERT INTO user_entities (name, email, password) VALUES (?, ?, ?)`, user.Name, user.Email, password)
	}

	return nil
}

func NewUserSeeder() Seeder {
	return &UserSeeder{}
}
