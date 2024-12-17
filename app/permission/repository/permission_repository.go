package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	model *gorm.DB
}

// FindByKey implements domain.PermissionRepository.
func (permission *PermissionRepository) FindByKey(ctx *gin.Context, key string, trashed bool) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(ctx)
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

func (permission *PermissionRepository) FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(ctx)
	permissionEntity := &domain.PermissionEntity{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	err := permission.model.Where("id = ?", id).First(&permissionEntity).Error

	return permissionEntity, err
}

func (permission *PermissionRepository) FindByNameAndKey(ctx *gin.Context, name string, key string) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(ctx)

	permissionEntity := &domain.PermissionEntity{}
	permission.model.First(&permissionEntity, "name = ? and key = ?", name, key)

	return permissionEntity, nil
}

func (permission *PermissionRepository) Delete(ctx *gin.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(ctx)
	permission.model.Delete(&domain.PermissionEntity{}, id)
}

func (permission *PermissionRepository) ForceDelete(ctx *gin.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(ctx)
	permissionEntity := &domain.PermissionEntity{}
	permission.model.Unscoped().Where("id = ?", id).Find(&permissionEntity)
	permission.model.Unscoped().Delete(&permissionEntity)
}

func (permission *PermissionRepository) Update(ctx *gin.Context, id uuid.UUID, payload *domain.PermissionEntity) error {
	permission.model = permission.model.WithContext(ctx)
	err := permission.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (permission *PermissionRepository) Create(ctx *gin.Context, payload *domain.PermissionEntity) error {
	permission.model = permission.model.WithContext(ctx)
	err := permission.model.Create(&payload).Error
	return err
}

func (permission *PermissionRepository) IsKeyExist(ctx *gin.Context, key string) bool {
	permission.model = permission.model.WithContext(ctx)
	var count int64
	permission.model.
		Where("key = ?", key).
		Count(&count)
	return count > 0
}

func (permission *PermissionRepository) IsKeyExistExceptPermissionId(ctx *gin.Context, key string, id uuid.UUID) bool {
	permission.model = permission.model.WithContext(ctx)
	var count int64
	permission.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count)

	return count > 0
}
