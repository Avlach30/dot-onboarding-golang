package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
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

func (role *RoleRepository) Pagination(httpContext *gin.Context) ([]domain.RoleEntity, int) {
	role.model = role.model.WithContext(httpContext)
	var roles []domain.RoleEntity
	var total int64

	// Query filter
	role.queryFilter(httpContext)
	// Query sort
	role.querySort(httpContext)

	role.model.Session(&gorm.Session{}).
		Scopes(utils.Paginate(httpContext)).
		Find(&roles).
		Count(&total)

	return roles, int(total)
}

// func filter for pagination
func (role *RoleRepository) queryFilter(httpContext *gin.Context) *gorm.DB {
	if search := httpContext.Query("search"); search != "" {
		role.model = role.model.
			Where("name LIKE ?", "%"+search+"%")
	}

	return role.model
}

// func query sort for pagination
func (role *RoleRepository) querySort(httpContext *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if sort := httpContext.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		role.model = role.model.
			Order(sort + " " + httpContext.Query("order"))

		// Handle order by asc or desc
		if order := httpContext.Query("order"); order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}

			role.model = role.model.
				Order(sort + " " + order)
		}
	}

	return role.model
}

// FindByKey implements domain.RoleRepository.
func (role *RoleRepository) FindByKey(httpContext *gin.Context, key string, trashed bool) (*domain.RoleEntity, error) {
	role.model = role.model.WithContext(httpContext)
	roleEntity := &domain.RoleEntity{}

	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.Where("key = ?", key).First(&roleEntity).Error

	return roleEntity, err
}

func (role *RoleRepository) FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) (*domain.RoleEntity, error) {
	role.model = role.model.WithContext(httpContext)
	roleEntity := &domain.RoleEntity{}
	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.
		Preload("Permissions").
		First(&roleEntity, id).
		Error

	return roleEntity, err
}

func (role *RoleRepository) FindByNameAndKey(httpContext *gin.Context, name string, key string) (*domain.RoleEntity, error) {
	role.model = role.model.WithContext(httpContext)
	roleEntity := &domain.RoleEntity{}
	role.model.First(&roleEntity, "name = ? and key = ?", name, key)

	return roleEntity, nil
}

func (role *RoleRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	role.model = role.model.WithContext(httpContext)
	role.model.Delete(&domain.RoleEntity{}, id)
}

func (role *RoleRepository) ForceDelete(httpContext *gin.Context, id uuid.UUID) {
	role.model = role.model.WithContext(httpContext)
	roleEntity := &domain.RoleEntity{}
	role.model.Unscoped().Delete(&roleEntity, id)
}

func (role *RoleRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.RoleEntity) error {
	role.model = role.model.WithContext(httpContext)
	err := role.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (role *RoleRepository) Create(httpContext *gin.Context, payload *domain.RoleEntity) error {
	role.model = role.model.WithContext(httpContext)
	err := role.model.Create(&payload).Error
	return err
}

func (role *RoleRepository) IsKeyExist(httpContext *gin.Context, key string) bool {
	role.model = role.model.WithContext(httpContext)
	var count int64
	role.model.
		Where("key = ?", key).
		Count(&count)
	return count > 0
}

func (role *RoleRepository) IsKeyExistExceptRoleId(httpContext *gin.Context, key string, id uuid.UUID) bool {
	role.model = role.model.WithContext(httpContext)
	var count int64
	role.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count)

	return count > 0
}
