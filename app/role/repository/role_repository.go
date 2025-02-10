package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type RoleRepository struct {
	model *gorm.DB
}

func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &RoleRepository{
		model: db.Model(&entities.RoleEntity{}),
	}
}

// Pagination get role data with pagination
func (role *RoleRepository) Pagination(ctx *gin.Context) ([]entities.RoleEntity, int) {
	query := role.model.WithContext(ctx)
	var roles []entities.RoleEntity
	var total int64

	// Query filter
	query = role.queryFilter(query, ctx)
	// Query sort
	query = role.querySort(query, ctx)

	err := query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(ctx)).
		Find(&roles).
		Count(&total).Error

	if err != nil {
		log.Println("Error pagination role", err)
		panic(*exception.ServerErrorException(err))
	}

	return roles, int(total)
}

// func filter for pagination
func (role *RoleRepository) queryFilter(query *gorm.DB, ctx *gin.Context) *gorm.DB {
	if search := ctx.Query("search"); search != "" {
		query = query.Where("name LIKE ?", search+"%")
	}

	return query
}

// func query sort for pagination
func (role *RoleRepository) querySort(query *gorm.DB, ctx *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if sort := ctx.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		// Handle order query
		if order := ctx.Query("order"); order != "" {
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

func (role *RoleRepository) FindOneById(ctx *gin.Context, id uuid.UUID, trashed bool) *entities.RoleEntity {
	role.model = role.model.WithContext(ctx)
	roleEntity := &entities.RoleEntity{}
	if trashed {
		role.model = role.model.Unscoped()
	}

	err := role.model.
		Preload("Permissions").
		First(&roleEntity, id).
		Error

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Role not found"))
	} else if err != nil {
		log.Println("Error role find by id: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return roleEntity
}

func (role *RoleRepository) Delete(ctx *gin.Context, id uuid.UUID) {
	role.model = role.model.WithContext(ctx)
	err := role.model.Delete(&entities.RoleEntity{}, id).Error

	if err != nil {
		log.Println("Error role delete: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (role *RoleRepository) Update(ctx *gin.Context, id uuid.UUID, payload *entities.RoleEntity) {
	role.model = role.model.WithContext(ctx)
	err := role.model.Where("id = ?", id).Updates(&payload).Error

	if err != nil {
		log.Println("Error role update: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (role *RoleRepository) Create(ctx *gin.Context, payload *entities.RoleEntity) {
	role.model = role.model.WithContext(ctx)
	err := role.model.Create(&payload).Error

	if err != nil {
		log.Println("Error role create: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (role *RoleRepository) IsKeyExist(ctx *gin.Context, key string) bool {
	role.model = role.model.WithContext(ctx)
	var count int64
	err := role.model.
		Where("key = ?", key).
		Count(&count).Error

	if err != nil {
		log.Println("Error role is key exist: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return count > 0
}

func (role *RoleRepository) IsKeyExistExceptRoleId(ctx *gin.Context, key string, id uuid.UUID) bool {
	role.model = role.model.WithContext(ctx)
	var count int64
	err := role.model.
		Where("key = ? AND id != ?", key, id).
		Count(&count).Error

	if err != nil {
		log.Println("Error role is key exist except role id: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return count > 0
}
