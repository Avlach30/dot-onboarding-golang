package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
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
func (role *RoleRepository) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.RoleEntity, int) {
	query := role.model.WithContext(httpContext)
	var roles []entities.RoleEntity
	var total int64

	// Query filter
	query = role.queryFilter(query, queryDto)
	// Query sort
	query = role.querySort(query, queryDto)

	// Count all column first before paginate the query
	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count user", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(queryDto)).
		Find(&roles).Error

	if err != nil {
		log.Println("Error pagination role", err)
		panic(*exception.ServerErrorException(err))
	}

	return roles, int(total)
}

// func filter for pagination
func (role *RoleRepository) queryFilter(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	if search := queryDto.Search; search != "" {
		query = query.Where("name ILIKE ?", search+"%")
	}

	return query
}

// func query sort for pagination
func (role *RoleRepository) querySort(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	sortableColumns := []string{"name", "created_at", "updated_at"}

	if sort := queryDto.SortBy; sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		// Handle order query
		if order := queryDto.Order; order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}
			query = query.Order(sort + " " + order)
		}
	} else {
		query = query.Order("updated_at desc")
	}

	return query
}

func (role *RoleRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.RoleEntity {
	role.model = role.model.WithContext(httpContext)
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

func (role *RoleRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	role.model = role.model.WithContext(httpContext)
	err := role.model.Delete(&entities.RoleEntity{}, id).Error

	if err != nil {
		log.Println("Error role delete: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (role *RoleRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.RoleEntity) {
	role.model = role.model.WithContext(httpContext)
	err := role.model.Where("id = ?", id).Updates(&payload).Error

	if err != nil {
		log.Println("Error role update: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (role *RoleRepository) Create(httpContext *gin.Context, payload *entities.RoleEntity) {
	role.model = role.model.WithContext(httpContext)
	err := role.model.Create(&payload).Error

	if err != nil {
		log.Println("Error role create: ", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (role *RoleRepository) IsKeyExist(httpContext *gin.Context, key string) bool {
	role.model = role.model.WithContext(httpContext)
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

func (role *RoleRepository) IsKeyExistExceptRoleId(httpContext *gin.Context, key string, id uuid.UUID) bool {
	role.model = role.model.WithContext(httpContext)
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
