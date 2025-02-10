package repository

import (
	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	permissionEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/entities"
	roleEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/entities"
	userEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/entities"
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
		userModel:           db.Model(&userEntities.UserEntity{}),
		permissionModel:     db.Model(&permissionEntities.PermissionEntity{}),
		roleModel:           db.Model(&roleEntities.RoleEntity{}),
		rolePermissionModel: db.Model(&roleEntities.RolePermissionEntity{}),
		userRoleModel:       db.Model(&userEntities.UserRoleEntity{}),
	}
}

// FindUserByEmail implements domain.AuthRepository.
func (authRepo *AuthRepository) FindUserByEmailWithRoles(httpContext *gin.Context, email string) (*userEntities.UserEntity, error) {
	authRepo.userModel = authRepo.userModel.WithContext(httpContext)
	user := &userEntities.UserEntity{}
	err := authRepo.userModel.
		Preload("Roles").
		Preload("Roles.Permissions").
		Where("email = ?", email).
		Find(&user).Error
	return user, err
}
