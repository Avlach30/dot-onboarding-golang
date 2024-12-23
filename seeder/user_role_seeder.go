package seeder

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	domainRole "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	domainUser "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gorm.io/gorm"
)

type UserRoleSeeder struct {
	RoleKey string `json:"role_key"`
	Email   string `json:"email"`
}

// Handle implements Seeder.
func (userRoleSeeder *UserRoleSeeder) Handle(db *gorm.DB) error {
	filePath := "seeder/files/user_roles.json"

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open JSON file: %v", err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("could not read JSON file: %v", err)
	}

	// Parse JSON data
	var userRoles []UserRoleSeeder
	if err := json.Unmarshal(bytes, &userRoles); err != nil {
		return fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	// Insert each user into the database
	db.Exec(`DELETE FROM user_roles`)
	for _, userRole := range userRoles {
		role := &domainRole.RoleEntity{}
		errorFindRole := db.Where("key = ?", userRole.RoleKey).First(role)
		if errorFindRole.Error != nil {
			log.Println("Error querying role:", errorFindRole.Error)
			continue
		}

		user := &domainUser.UserEntity{}
		errorFindUser := db.Where("email = ?", userRole.Email).First(user)
		if errorFindUser.Error != nil {
			log.Println("Error querying user:", errorFindUser.Error)
			continue
		}

		if role.ID == uuid.Nil || user.ID == uuid.Nil {
			log.Println("role and/or user nil because : ", role.ID, user.ID)
			continue
		}

		log.Println(userRole)
		db.Exec(`INSERT INTO user_roles (role_id, user_id) VALUES (?, ?)`, role.ID, user.ID)
	}

	return nil
}

func NewUserRoleSeeder() Seeder {
	return &UserRoleSeeder{}
}
