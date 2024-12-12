package repository

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	model *gorm.DB
}

// FindByKey implements domain.PermissionRepository.
func (permission *PermissionRepository) FindByKey(key string, trashed bool) (*domain.PermissionEntity, error) {
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

func (permission *PermissionRepository) FindById(id uuid.UUID, trashed bool) (*domain.PermissionEntity, error) {
	permissionEntity := &domain.PermissionEntity{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	permission.model.Where("id = ?", id).First(&permissionEntity)

	return permissionEntity, nil
}

func (permission *PermissionRepository) FindByNameAndKey(name string, key string) (*domain.PermissionEntity, error) {

	permissionEntity := &domain.PermissionEntity{}
	permission.model.First(&permissionEntity, "name = ? and key = ?", name, key)

	return permissionEntity, nil
}

func (permission *PermissionRepository) Delete(id uuid.UUID) {
	permission.model.Where("id = ?", id).Delete(&domain.PermissionEntity{})
}

func (permission *PermissionRepository) ForceDelete(id uuid.UUID) {
	permissionEntity := &domain.PermissionEntity{}
	permission.model.Unscoped().Where("id = ?", id).Find(&permissionEntity)
	permission.model.Unscoped().Delete(&permissionEntity)
}

func (permission *PermissionRepository) Update(id uuid.UUID, payload *domain.PermissionEntity) {
	permission.model.Where("id = ?", id).Updates(&payload)
}

func (permission *PermissionRepository) Create(payload *domain.PermissionEntity) error {
	permission.model.Create(payload)

	return nil
}
