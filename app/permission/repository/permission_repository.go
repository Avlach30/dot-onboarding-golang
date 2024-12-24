package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	model *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) domain.PermissionRepository {
	return &PermissionRepository{
		model: db.Model(&domain.PermissionEntity{}),
	}
}

func (permission *PermissionRepository) Pagination(httpContext *gin.Context) ([]domain.PermissionEntity, int) {
	permission.model = permission.model.WithContext(httpContext)
	var permissions []domain.PermissionEntity
	var total int64

	// Query filter
	permission.queryFilter(httpContext)
	// Query sort
	permission.querySort(httpContext)

	permission.model.Session(&gorm.Session{}).
		Scopes(utils.Paginate(httpContext)).
		Find(&permissions).
		Count(&total)

	return permissions, int(total)
}

// func filter for pagination
func (permission *PermissionRepository) queryFilter(httpContext *gin.Context) *gorm.DB {
	if search := httpContext.Query("search"); search != "" {
		permission.model = permission.model.
			Where("name LIKE ?", "%"+search+"%")
	}

	return permission.model
}

// func query sort for pagination
func (permission *PermissionRepository) querySort(httpContext *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if sort := httpContext.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		permission.model = permission.model.
			Order(sort + " " + httpContext.Query("order"))

		// Handle order by asc or desc
		if order := httpContext.Query("order"); order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}

			permission.model = permission.model.
				Order(sort + " " + order)
		}
	}

	return permission.model
}

// FindByKey implements domain.PermissionRepository.
func (permission *PermissionRepository) FindByKey(httpContext *gin.Context, key string, trashed bool) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(httpContext)
	permissionEntity := &domain.PermissionEntity{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	permission.model.Where("key = ?", key).First(&permissionEntity)

	return permissionEntity, nil
}

func (permission *PermissionRepository) FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(httpContext)
	permissionEntity := &domain.PermissionEntity{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	err := permission.model.First(&permissionEntity, id).Error

	return permissionEntity, err
}

func (permission *PermissionRepository) FindByNameAndKey(httpContext *gin.Context, name string, key string) (*domain.PermissionEntity, error) {
	permission.model = permission.model.WithContext(httpContext)

	permissionEntity := &domain.PermissionEntity{}
	permission.model.First(&permissionEntity, "name = ? and key = ?", name, key)

	return permissionEntity, nil
}

func (permission *PermissionRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(httpContext)
	permission.model.Delete(&domain.PermissionEntity{}, id)
}

func (permission *PermissionRepository) ForceDelete(httpContext *gin.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(httpContext)
	permissionEntity := &domain.PermissionEntity{}
	permission.model.Unscoped().Where("id = ?", id).Find(&permissionEntity)
	permission.model.Unscoped().Delete(&permissionEntity)
}

func (permission *PermissionRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.PermissionEntity) error {
	permission.model = permission.model.WithContext(httpContext)
	err := permission.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (permission *PermissionRepository) Create(httpContext *gin.Context, payload *domain.PermissionEntity) error {
	permission.model = permission.model.WithContext(httpContext)
	err := permission.model.Create(&payload).Error
	return err
}

func (permission *PermissionRepository) IsKeyExist(httpContext *gin.Context, key string) bool {
	permission.model = permission.model.WithContext(httpContext)
	var count int64
	permission.model.
		Where("key = ?", key).
		Count(&count)
	return count > 0
}

func (permission *PermissionRepository) IsKeyExistExceptPermissionId(httpContext *gin.Context, key string, id uuid.UUID) bool {
	permission.model = permission.model.WithContext(httpContext)
	var count int64
	permission.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count)

	return count > 0
}
