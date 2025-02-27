package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	permissionEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	roleEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	userEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
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

func (authRepo *AuthRepository) FindUserByIDWithRoles(httpContext *gin.Context, id uuid.UUID) (*userEntities.UserEntity, error) {
	authRepo.userModel = authRepo.userModel.WithContext(httpContext)
	user := &userEntities.UserEntity{}
	err := authRepo.userModel.
		Preload("Roles").
		Preload("Roles.Permissions").
		Where("id = ?", id).
		Find(&user).Error

	return user, err
}
