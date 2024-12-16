package repository

import (
	"context"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gorm.io/gorm"
)

type RoleRepository struct {
	model *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &RoleRepository{
		model: db.Model(&domain.RoleEntity{}),
	}
}

// FindByKey implements domain.RoleRepository.
func (role *RoleRepository) FindByKey(context *context.Context, key string, trashed bool) (*domain.RoleEntity, error) {
	role.model = role.model.WithContext(*context)
	roleEntity := &domain.RoleEntity{}

	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.Where("key = ?", key).First(&roleEntity).Error

	return roleEntity, err
}

func (role *RoleRepository) FindById(context *context.Context, id uuid.UUID, trashed bool) (*domain.RoleEntity, error) {
	role.model = role.model.WithContext(*context)
	roleEntity := &domain.RoleEntity{}
	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.Where("id = ?", id).First(&roleEntity).Error

	return roleEntity, err
}

func (role *RoleRepository) FindByNameAndKey(context *context.Context, name string, key string) (*domain.RoleEntity, error) {
	role.model = role.model.WithContext(*context)
	roleEntity := &domain.RoleEntity{}
	role.model.First(&roleEntity, "name = ? and key = ?", name, key)

	return roleEntity, nil
}

func (role *RoleRepository) Delete(context *context.Context, id uuid.UUID) {
	role.model = role.model.WithContext(*context)
	role.model.Delete(&domain.RoleEntity{}, id)
}

func (role *RoleRepository) ForceDelete(context *context.Context, id uuid.UUID) {
	role.model = role.model.WithContext(*context)
	roleEntity := &domain.RoleEntity{}
	role.model.Unscoped().Delete(&roleEntity, id)
}

func (role *RoleRepository) Update(context *context.Context, id uuid.UUID, payload *domain.RoleEntity) error {
	role.model = role.model.WithContext(*context)
	err := role.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (role *RoleRepository) Create(context *context.Context, payload *domain.RoleEntity) error {
	role.model = role.model.WithContext(*context)
	err := role.model.Create(&payload).Error
	return err
}

func (role *RoleRepository) IsKeyExist(context *context.Context, key string) bool {
	role.model = role.model.WithContext(*context)
	var count int64
	role.model.
		Where("key = ?", key).
		Count(&count)
	return count > 0
}

func (role *RoleRepository) IsKeyExistExceptRoleId(context *context.Context, key string, id uuid.UUID) bool {
	role.model = role.model.WithContext(*context)
	var count int64
	role.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count)

	return count > 0
}
