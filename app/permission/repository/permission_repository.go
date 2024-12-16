package repository

import (
	"context"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	model *gorm.DB
}

// FindByKey implements domain.PermissionRepository.
func (permission *PermissionRepository) FindByKey(context *context.Context, key string, trashed bool) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(*context)
	permissionEntity := &domain.PermissionEntity{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	permission.model.Where("key = ?", key).First(&permissionEntity)

	return permissionEntity, nil
}

func NewPermissionRepository(db *gorm.DB) domain.PermissionRepository {
	return &PermissionRepository{
		model: db.Model(&domain.PermissionEntity{}),
	}
}

func (permission *PermissionRepository) FindById(context *context.Context, id uuid.UUID, trashed bool) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(*context)
	permissionEntity := &domain.PermissionEntity{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	err := permission.model.Where("id = ?", id).First(&permissionEntity).Error

	return permissionEntity, err
}

func (permission *PermissionRepository) FindByNameAndKey(context *context.Context, name string, key string) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(*context)

	permissionEntity := &domain.PermissionEntity{}
	permission.model.First(&permissionEntity, "name = ? and key = ?", name, key)

	return permissionEntity, nil
}

func (permission *PermissionRepository) Delete(context *context.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(*context)
	permission.model.Delete(&domain.PermissionEntity{}, id)
}

func (permission *PermissionRepository) ForceDelete(context *context.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(*context)
	permissionEntity := &domain.PermissionEntity{}
	permission.model.Unscoped().Where("id = ?", id).Find(&permissionEntity)
	permission.model.Unscoped().Delete(&permissionEntity)
}

func (permission *PermissionRepository) Update(context *context.Context, id uuid.UUID, payload *domain.PermissionEntity) error {
	permission.model = permission.model.WithContext(*context)
	err := permission.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (permission *PermissionRepository) Create(context *context.Context, payload *domain.PermissionEntity) error {
	permission.model = permission.model.WithContext(*context)
	err := permission.model.Create(&payload).Error
	return err
}

func (permission *PermissionRepository) IsKeyExist(context *context.Context, key string) bool {
	permission.model = permission.model.WithContext(*context)
	var count int64
	permission.model.
		Where("key = ?", key).
		Count(&count)
	return count > 0
}

func (permission *PermissionRepository) IsKeyExistExceptPermissionId(context *context.Context, key string, id uuid.UUID) bool {
	permission.model = permission.model.WithContext(*context)
	var count int64
	permission.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count)

	return count > 0
}
