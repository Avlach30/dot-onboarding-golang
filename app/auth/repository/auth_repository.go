package repository

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	permissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	rolePermissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role_permission/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gorm.io/gorm"
)

type AuthRepository struct {
	userModel           *gorm.DB
	permissionModel     *gorm.DB
	roleModel           *gorm.DB
	rolePermissionModel *gorm.DB
}

// FindUserByEmail implements domain.AuthRepository.
func (authRepo *AuthRepository) FindUserByEmail(email string) (*userDomain.UserEntity, error) {
	user := &userDomain.UserEntity{}
	err := authRepo.userModel.Where("email = ?", email).Find(&user).Error
	return user, err
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &AuthRepository{
		userModel:           db.Model(&userDomain.UserEntity{}),
		permissionModel:     db.Model(&permissionDomain.PermissionEntity{}),
		roleModel:           db.Model(&roleDomain.RoleEntity{}),
		rolePermissionModel: db.Model(&rolePermissionDomain.RolePermissionEntity{}),
	}
}
