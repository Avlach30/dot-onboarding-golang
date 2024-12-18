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
		model: db.Model(&domain.Permission{}),
	}
}

func (permission *PermissionRepository) Pagination(ctx *gin.Context) ([]domain.Permission, int) {
	permission.model = permission.model.WithContext(ctx)
	var permissions []domain.Permission
	var total int64

	// Query filter
	permission.queryFilter(ctx)
	// Query sort
	permission.querySort(ctx)

	permission.model.Session(&gorm.Session{}).
		Scopes(utils.Paginate(ctx)).
		Find(&permissions).
		Count(&total)

	return permissions, int(total)
}

// func filter for pagination
func (permission *PermissionRepository) queryFilter(ctx *gin.Context) *gorm.DB {
	if search := ctx.Query("search"); search != "" {
		permission.model = permission.model.
			Where("name LIKE ?", "%"+search+"%")
	}

	return permission.model
}

// func query sort for pagination
func (permission *PermissionRepository) querySort(ctx *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if sort := ctx.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		permission.model = permission.model.
			Order(sort + " " + ctx.Query("order"))

		// Handle order by asc or desc
		if order := ctx.Query("order"); order != "" {
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
func (permission *PermissionRepository) FindByKey(ctx *gin.Context, key string, trashed bool) (*domain.Permission, error) {
	permission.model = permission.model.WithContext(ctx)
	Permission := &domain.Permission{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	permission.model.Where("key = ?", key).First(&Permission)

	return Permission, nil
}

func (permission *PermissionRepository) FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*domain.Permission, error) {
	permission.model = permission.model.WithContext(ctx)
	Permission := &domain.Permission{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	err := permission.model.Where("id = ?", id).First(&Permission).Error

	return Permission, err
}

func (permission *PermissionRepository) FindByNameAndKey(ctx *gin.Context, name string, key string) (*domain.Permission, error) {
	permission.model = permission.model.WithContext(ctx)

	Permission := &domain.Permission{}
	permission.model.First(&Permission, "name = ? and key = ?", name, key)

	return Permission, nil
}

func (permission *PermissionRepository) Delete(ctx *gin.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(ctx)
	permission.model.Delete(&domain.Permission{}, id)
}

func (permission *PermissionRepository) ForceDelete(ctx *gin.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(ctx)
	Permission := &domain.Permission{}
	permission.model.Unscoped().Where("id = ?", id).Find(&Permission)
	permission.model.Unscoped().Delete(&Permission)
}

func (permission *PermissionRepository) Update(ctx *gin.Context, id uuid.UUID, payload *domain.Permission) error {
	permission.model = permission.model.WithContext(ctx)
	err := permission.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (permission *PermissionRepository) Create(ctx *gin.Context, payload *domain.Permission) error {
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
