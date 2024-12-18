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
		model: db.Model(&domain.Role{}),
	}
}

func (role *RoleRepository) Pagination(ctx *gin.Context) ([]domain.Role, int) {
	role.model = role.model.WithContext(ctx)
	var roles []domain.Role
	var total int64

	// Query filter
	role.queryFilter(ctx)
	// Query sort
	role.querySort(ctx)

	role.model.Session(&gorm.Session{}).
		Scopes(utils.Paginate(ctx)).
		Find(&roles).
		Count(&total)

	return roles, int(total)
}

// func filter for pagination
func (role *RoleRepository) queryFilter(ctx *gin.Context) *gorm.DB {
	if search := ctx.Query("search"); search != "" {
		role.model = role.model.
			Where("name LIKE ?", "%"+search+"%")
	}

	return role.model
}

// func query sort for pagination
func (role *RoleRepository) querySort(ctx *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if sort := ctx.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		role.model = role.model.
			Order(sort + " " + ctx.Query("order"))

		// Handle order by asc or desc
		if order := ctx.Query("order"); order != "" {
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
func (role *RoleRepository) FindByKey(ctx *gin.Context, key string, trashed bool) (*domain.Role, error) {
	role.model = role.model.WithContext(ctx)
	Role := &domain.Role{}

	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.Where("key = ?", key).First(&Role).Error

	return Role, err
}

func (role *RoleRepository) FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*domain.Role, error) {
	role.model = role.model.WithContext(ctx)
	Role := &domain.Role{}
	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.
		Preload("Permissions").
		Where("id = ?", id).
		First(&Role).
		Error

	return Role, err
}

func (role *RoleRepository) FindByNameAndKey(ctx *gin.Context, name string, key string) (*domain.Role, error) {
	role.model = role.model.WithContext(ctx)
	Role := &domain.Role{}
	role.model.First(&Role, "name = ? and key = ?", name, key)

	return Role, nil
}

func (role *RoleRepository) Delete(ctx *gin.Context, id uuid.UUID) {
	role.model = role.model.WithContext(ctx)
	role.model.Delete(&domain.Role{}, id)
}

func (role *RoleRepository) ForceDelete(ctx *gin.Context, id uuid.UUID) {
	role.model = role.model.WithContext(ctx)
	Role := &domain.Role{}
	role.model.Unscoped().Delete(&Role, id)
}

func (role *RoleRepository) Update(ctx *gin.Context, id uuid.UUID, payload *domain.Role) error {
	role.model = role.model.WithContext(ctx)
	err := role.model.Where("id = ?", id).Updates(&payload).Error
	return err
}

func (role *RoleRepository) Create(ctx *gin.Context, payload *domain.Role) error {
	role.model = role.model.WithContext(ctx)
	err := role.model.Create(&payload).Error
	return err
}

func (role *RoleRepository) IsKeyExist(ctx *gin.Context, key string) bool {
	role.model = role.model.WithContext(ctx)
	var count int64
	role.model.
		Where("key = ?", key).
		Count(&count)
	return count > 0
}

func (role *RoleRepository) IsKeyExistExceptRoleId(ctx *gin.Context, key string, id uuid.UUID) bool {
	role.model = role.model.WithContext(ctx)
	var count int64
	role.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count)

	return count > 0
}
