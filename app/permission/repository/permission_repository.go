package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type PermissionRepository struct {
	model *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) domain.PermissionRepository {
	return &PermissionRepository{
		model: db.Model(&entities.PermissionEntity{}),
	}
}

// Pagination get permission data with pagination
func (permission *PermissionRepository) Pagination(httpContext *gin.Context) ([]entities.PermissionEntity, int) {
	query := permission.model.WithContext(httpContext)
	var permissions []entities.PermissionEntity
	var total int64

	// Query filter
	query = permission.queryFilter(query, httpContext)
	// Query sort
	query = permission.querySort(query, httpContext)

	// Count all column first before paginate the query
	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count user", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(httpContext)).Find(&permissions).Error

	if err != nil {
		log.Println("Error pagination permission", err)
		panic(*exception.ServerErrorException(err))
	}

	return permissions, int(total)
}

// func filter for pagination
func (permission *PermissionRepository) queryFilter(query *gorm.DB, httpContext *gin.Context) *gorm.DB {
	if search := httpContext.Query("search"); search != "" {
		query = query.Where("name LIKE ?", search+"%")
	}

	return query
}

// func query sort for pagination
func (permission *PermissionRepository) querySort(query *gorm.DB, httpContext *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if sort := httpContext.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		// Handle order query
		if order := httpContext.Query("order"); order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}
			query = query.Order(sort + " " + order)
		} else {
			query = query.Order(sort)
		}
	}

	return query
}

func (permission *PermissionRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.PermissionEntity {
	permission.model = permission.model.WithContext(httpContext)
	permissionEntity := &entities.PermissionEntity{}
	if trashed {
		permission.model = permission.model.Unscoped()
	}

	err := permission.model.First(&permissionEntity, id).Error
	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Permission not found"))
	} else if err != nil {
		log.Println("Error find permission by id", err)
		panic(*exception.ServerErrorException(err))
	}

	return permissionEntity
}

func (permission *PermissionRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	permission.model = permission.model.WithContext(httpContext)
	err := permission.model.Delete(&entities.PermissionEntity{}, id).Error

	if err != nil {
		log.Println("Error delete permission", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (permission *PermissionRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.PermissionEntity) {
	permission.model = permission.model.WithContext(httpContext)
	err := permission.model.Where("id = ?", id).Updates(&payload).Error
	if err != nil {
		log.Println("Error update permission", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (permission *PermissionRepository) Create(httpContext *gin.Context, payload *entities.PermissionEntity) {
	permission.model = permission.model.WithContext(httpContext)
	err := permission.model.Create(&payload).Error
	if err != nil {
		log.Println("Error create permission", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (permission *PermissionRepository) IsKeyExist(httpContext *gin.Context, key string) bool {
	permission.model = permission.model.WithContext(httpContext)
	var count int64
	err := permission.model.
		Where("key = ?", key).
		Count(&count).Error

	if err != nil {
		log.Println("Error check key exist", err)
		panic(*exception.ServerErrorException(err))
	}

	return count > 0
}

func (permission *PermissionRepository) IsKeyExistExceptPermissionId(httpContext *gin.Context, key string, id uuid.UUID) bool {
	permission.model = permission.model.WithContext(httpContext)
	var count int64
	err := permission.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count).Error

	if err != nil {
		log.Println("Error check key exist", err)
		panic(*exception.ServerErrorException(err))
	}

	return count > 0
}
