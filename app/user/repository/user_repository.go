package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	roleEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	model     *gorm.DB
	roleModel *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		model:     db.Model(&entities.UserEntity{}),
		roleModel: db.Model(&roleEntities.RoleEntity{}),
	}
}

// Pagination get user data with pagination
func (user *UserRepository) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.UserEntity, int) {
	query := user.model.WithContext(httpContext)
	var users []entities.UserEntity
	var total int64

	// Query filter
	query = user.queryFilter(query, queryDto)

	// Query sort
	query = user.querySort(query, queryDto)

	// Count all column first before paginate the query
	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error count user", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(queryDto)).
		Find(&users).Error

	if err != nil {
		log.Println("Error pagination user", err)
		panic(*exception.ServerErrorException(err))
	}

	return users, int(total)
}

// func filter for pagination
func (user *UserRepository) queryFilter(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	if search := queryDto.Search; search != "" {
		query = query.Where("name ILIKE ?", search+"%")
	}

	return query
}

// func query sort for pagination
func (user *UserRepository) querySort(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	sortableColumns := []string{"name", "email", "created_at", "updated_at"}

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

func (user *UserRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.UserEntity {
	user.model = user.model.WithContext(httpContext)
	userEntity := &entities.UserEntity{}
	if trashed {
		user.model = user.model.Unscoped()
	}

	err := user.model.
		Preload("Roles").
		Preload("Roles.Permissions").
		First(&userEntity, id).
		Error

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("User not found"))
	} else if err != nil {
		log.Println("Error find user by id", err)
		panic(*exception.ServerErrorException(err))
	}

	return userEntity
}

func (user *UserRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	user.model = user.model.WithContext(httpContext)
	err := user.model.Delete(&entities.UserEntity{}, id).Error

	if err != nil {
		log.Println("Error delete user", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (user *UserRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.UserEntity) {
	user.model = user.model.WithContext(httpContext)
	userEntity := &entities.UserEntity{}

	err := user.model.Transaction(func(tx *gorm.DB) error {
		// Find user within transaction
		if err := tx.First(&userEntity, id).Error; err != nil {
			return err
		}

		// Permanently delete user roles relationships
		if err := tx.Unscoped().Model(&userEntity).
			Association("Roles").Unscoped().Clear(); err != nil {
			return err
		}

		// Update user data within transaction
		if err := tx.Where("id = ?", id).Updates(&payload).Error; err != nil {
			return err
		}

		// Create user roles
		if err := tx.Model(&userEntity).Association("Roles").Append(payload.Roles); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Error update user", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (user *UserRepository) Create(httpContext *gin.Context, payload *entities.UserEntity) {
	user.model = user.model.WithContext(httpContext)
	err := user.model.Create(&payload).Error

	if err != nil {
		log.Println("Error create user", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (user *UserRepository) IsEmailExist(httpContext *gin.Context, email string) bool {
	user.model = user.model.WithContext(httpContext)
	var count int64
	user.model.
		Where("email = ?", email).
		Count(&count)

	return count > 0
}

func (user *UserRepository) IsEmailExistExceptUserId(httpContext *gin.Context, email string, id uuid.UUID) bool {
	user.model = user.model.WithContext(httpContext)
	var count int64
	err := user.model.
		Where("email = ? AND id != ?", email, id).
		Count(&count).Error

	if err != nil {
		log.Println("Error checking email exist", err)
		panic(*exception.BussinessException(err.Error()))
	}

	return count > 0
}

func (user *UserRepository) FindRoleByIds(httpContext *gin.Context, ids []uuid.UUID) []roleEntities.RoleEntity {
	user.roleModel = user.roleModel.WithContext(httpContext)
	var roleEntities []roleEntities.RoleEntity
	err := user.roleModel.Where("id IN ?", ids).Find(&roleEntities).Error

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Roles not found"))
	} else if err != nil {
		log.Println("Error finding roles", err)
		panic(*exception.ServerErrorException(err))
	}

	if len(roleEntities) == 0 {
		panic(*exception.NotFoundException("Roles not found"))
	}

	return roleEntities
}

// Delete user's roles using association table without deleting the user
func (user *UserRepository) DeleteUserRoles(httpContext *gin.Context, id uuid.UUID) {
	user.model = user.model.WithContext(httpContext)
	// Permanently delete user roles relationships
	err := user.model.Unscoped().Association("Roles").Unscoped().Clear()

	if err != nil {
		log.Println("Error delete user roles", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (user *UserRepository) IsExistById(httpContext *gin.Context, id uuid.UUID) bool {
	user.model = user.model.WithContext(httpContext)

	var userEntity entities.UserEntity
	err := user.model.
		First(&userEntity, id).
		Error

	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Println("Error fetching user by id", err)
		panic(*exception.ServerErrorException(err))
	}

	return true
}
