package repository

import (
	"context"

	"github.com/google/uuid"
	permissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	rolePermissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role_permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	userModel           *gorm.DB
	permissionModel     *gorm.DB
	roleModel           *gorm.DB
	rolePermissionModel *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		userModel:           db.Model(&userDomain.UserEntity{}),
		permissionModel:     db.Model(&permissionDomain.PermissionEntity{}),
		roleModel:           db.Model(&roleDomain.RoleEntity{}),
		rolePermissionModel: db.Model(&rolePermissionDomain.RolePermissionEntity{}),
	}
}

func (user *UserRepository) FindById(context *context.Context, id uuid.UUID, trashed bool) (*domain.UserEntity, error) {
	user.userModel = user.userModel.WithContext(*context)
	userEntity := &domain.UserEntity{}
	if trashed {
		user.userModel = user.userModel.Unscoped()
	}

	err := user.userModel.Where("id = ?", id).First(&userEntity).Error

	return userEntity, err
}

func (user *UserRepository) Delete(context *context.Context, id uuid.UUID) {
	user.userModel = user.userModel.WithContext(*context)
	user.userModel.Delete(&domain.UserEntity{}, id)
}

func (user *UserRepository) ForceDelete(context *context.Context, id uuid.UUID) {
	user.userModel = user.userModel.WithContext(*context)
	userEntity := &domain.UserEntity{}
	user.userModel.Unscoped().Delete(&userEntity, id)
}

func (user *UserRepository) Update(context *context.Context, id uuid.UUID, payload *domain.UserEntity) error {
	user.userModel = user.userModel.WithContext(*context)
	err := user.userModel.Where("id = ?", id).Updates(&payload).Error

	return err
}

func (user *UserRepository) Create(context *context.Context, payload *domain.UserEntity) error {
	user.userModel = user.userModel.WithContext(*context)
	err := user.userModel.Create(&payload).Error
	return err
}

func (user *UserRepository) IsEmailExist(context *context.Context, email string) bool {
	user.userModel = user.userModel.WithContext(*context)
	var count int64
	user.userModel.
		Where("email = ?", email).
		Count(&count)
	return count > 0
}

func (user *UserRepository) IsEmailExistExceptUserId(context *context.Context, email string, id uuid.UUID) bool {
	user.userModel = user.userModel.WithContext(*context)
	var count int64
	user.userModel.
		Where("email = ? AND id != ?", email, id).
		Count(&count)

	return count > 0
}
