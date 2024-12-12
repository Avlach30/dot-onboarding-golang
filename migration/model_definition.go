package migration

import (
	permissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	rolePermissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role_permission/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gorm.io/gorm"
)

func Run(db *gorm.DB, isAutoMigration bool) {

	entities := []interface{}{
		&userDomain.UserEntity{},
		&roleDomain.RoleEntity{},
		&permissionDomain.PermissionEntity{},
		&rolePermissionDomain.RolePermissionEntity{},
	}

	if isAutoMigration {
		AutoMigrate(db, entities)
	} else {
		// migration for existing project
		panic("unimplemented")
	}
}
