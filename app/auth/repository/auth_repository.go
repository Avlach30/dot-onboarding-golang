package repository

import (
	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	permissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	rolePermissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role_permission/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	userRoleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user_role/domain"
	"gorm.io/gorm"
)

type AuthRepository struct {
	userModel           *gorm.DB
	permissionModel     *gorm.DB
	roleModel           *gorm.DB
	userRoleModel       *gorm.DB
	rolePermissionModel *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &AuthRepository{
		userModel:           db.Model(&userDomain.UserEntity{}),
		permissionModel:     db.Model(&permissionDomain.PermissionEntity{}),
		roleModel:           db.Model(&roleDomain.RoleEntity{}),
		rolePermissionModel: db.Model(&rolePermissionDomain.RolePermissionEntity{}),
		userRoleModel:       db.Model(&userRoleDomain.UserRoleEntity{}),
	}
}

// FindUserByEmail implements domain.AuthRepository.
func (authRepo *AuthRepository) FindUserByEmailWithRoles(httpContext *gin.Context, email string) (*userDomain.UserEntity, error) {
	authRepo.userModel = authRepo.userModel.WithContext(httpContext)
	user := &userDomain.UserEntity{}
	err := authRepo.userModel.
		Preload("Roles").
		Preload("Roles.Permissions").
		Where("email = ?", email).
		Find(&user).Error
	return user, err
}
