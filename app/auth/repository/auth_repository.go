package repository

import (
	"context"

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
func (authRepo *AuthRepository) FindUserByEmail(context *context.Context, email string) (*userDomain.User, error) {
	authRepo.userModel = authRepo.userModel.WithContext(*context)
	user := &userDomain.User{}
	err := authRepo.userModel.Where("email = ?", email).Find(&user).Error
	return user, err
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &AuthRepository{
		userModel:           db.Model(&userDomain.User{}),
		permissionModel:     db.Model(&permissionDomain.Permission{}),
		roleModel:           db.Model(&roleDomain.Role{}),
		rolePermissionModel: db.Model(&rolePermissionDomain.RolePermission{}),
	}
}
