package seeder

import (
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, seederCommands []string) error {

	listSeeders := make(map[string]Seeder)

	listSeeders["UserSeeder"] = NewUserSeeder()
	listSeeders["PermissionSeeder"] = NewRoleSeeder()
	listSeeders["RoleSeeder"] = NewPermissionSeeder()

	if len(seederCommands) > 0 {
		for _, seederCommand := range seederCommands {
			err := listSeeders[seederCommand].Handle(db)
			if err != nil {
				return err
			}
		}
	} else {
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
