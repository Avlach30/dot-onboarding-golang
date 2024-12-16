package seeder

import (
	"log"

	"gorm.io/gorm"
)

func Run(db *gorm.DB, seederCommand []string) error {

	listSeeders := make(map[string]Seeder)

	listSeeders["UserSeeder"] = NewUserSeeder()
	listSeeders["PermissionSeeder"] = NewRoleSeeder()
	listSeeders["RoleSeeder"] = NewPermissionSeeder()


	if len(seederCommand) > 0 {
		for _, seederCommand := range seederCommand {
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
