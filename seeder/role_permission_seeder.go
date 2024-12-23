package seeder

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
	domainPermission "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	domainRole "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gorm.io/gorm"
)

type RolePermissionSeeder struct {
	RoleKey       string `json:"role_key"`
	PermissionKey string `json:"permission_key"`
}

// Handle implements Seeder.
func (userRolePermissionSeeder *RolePermissionSeeder) Handle(db *gorm.DB) error {
	filePath := "seeder/files/role_permissions.json"

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
	var rolePermissions []RolePermissionSeeder
	if err := json.Unmarshal(bytes, &rolePermissions); err != nil {
		return fmt.Errorf("could not unmarshal JSON data: %v", err)
	}

	// Insert each user into the database
	db.Exec(`DELETE FROM role_permissions`)
	for _, rolePermission := range rolePermissions {
		role := &domainRole.RoleEntity{}

		errorFindRole := db.Where("key = ?", rolePermission.RoleKey).First(role)
		if errorFindRole.Error != nil {
			log.Println("Error querying role:", errorFindRole.Error)
			continue
		}

		permission := &domainPermission.PermissionEntity{}
		errorFindPermission := db.Where("key = ?", rolePermission.PermissionKey).First(permission)
		if errorFindPermission.Error != nil {
			log.Println("Error querying permission:", errorFindPermission.Error)
			continue
		}

		if role.ID == uuid.Nil || permission.ID == uuid.Nil {
			log.Println("role and/or permission nil because : ", role.ID, permission.ID)
			continue
		}

		log.Println(rolePermission)
		db.Exec(`INSERT INTO role_permissions (role_id, permission_id) VALUES (?, ?)`, role.ID, permission.ID)
	}

	return nil
}

func NewRolePermissionSeeder() Seeder {
	return &RolePermissionSeeder{}
}
