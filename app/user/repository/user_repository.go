package repository

import (
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

func (user *UserRepository) FindById(id uuid.UUID, trashed bool) (*domain.UserEntity, error) {
	userEntity := &domain.UserEntity{}
	if trashed {
		user.userModel = user.userModel.Unscoped()
	}

	err := user.userModel.Where("id = ?", id).First(&userEntity).Error

	return userEntity, err
}

func (user *UserRepository) Delete(id uuid.UUID) {
	user.userModel.Where("id = ?", id).Delete(&domain.UserEntity{})
}

func (user *UserRepository) ForceDelete(id uuid.UUID) {

	userEntity := &domain.UserEntity{}
	user.userModel.Unscoped().Where("id = ?", id).Find(&userEntity)
	user.userModel.Unscoped().Delete(&userEntity)
}

func (user *UserRepository) Update(id uuid.UUID, payload *domain.UserEntity) {
	panic("unimplemented")
}

func (user *UserRepository) Create(payload *domain.UserEntity) error {
	return nil
}
