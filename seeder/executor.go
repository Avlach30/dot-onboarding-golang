package seeder

import (
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, seederCommands []string) error {

	if len(seederCommands) > 0 {

		listSeeders := make(map[string]Seeder)

		listSeeders["UserSeeder"] = NewUserSeeder()
		listSeeders["PermissionSeeder"] = NewRoleSeeder()
		listSeeders["RoleSeeder"] = NewPermissionSeeder()
		listSeeders["UserRoleSeeder"] = NewUserRoleSeeder()
		listSeeders["RolePermissionSeeder"] = NewRolePermissionSeeder()

		for _, seederCommand := range seederCommands {
			err := listSeeders[seederCommand].Handle(db)
			if err != nil {
				return err
			}
		}
	} else {
		db.Exec(`DELETE FROM role_permissions`)
		db.Exec(`DELETE FROM user_roles`)
		db.Exec(`DELETE FROM permissions`)
		db.Exec(`DELETE FROM roles`)
		db.Exec(`DELETE FROM users`)

		listSeeders := []Seeder{
			NewUserSeeder(),
			NewRoleSeeder(),
			NewPermissionSeeder(),
			NewUserRoleSeeder(),
			NewRolePermissionSeeder(),
		}

		for _, seeder := range listSeeders {
			err := seeder.Handle(db)
			if err != nil {
				return err
			}
		}
	}

	log.Println("Seed executed!")

	return nil
}
