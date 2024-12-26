package repository

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	model     *gorm.DB
	roleModel *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		model:     db.Model(&domain.UserEntity{}),
		roleModel: db.Model(&roleDomain.RoleEntity{}),
	}
}

// Pagination get user data with pagination
func (user *UserRepository) Pagination(httpContext *gin.Context) ([]domain.UserEntity, int) {
	query := user.model.WithContext(httpContext)
	var users []domain.UserEntity
	var total int64

	// Query filter
	query = user.queryFilter(query, httpContext)
	// Query sort
	query = user.querySort(query, httpContext)

	err := query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(httpContext)).
		Find(&users).
		Count(&total).Error

	if err != nil {
		log.Println("Error pagination user", err)
		panic(*exception.ServerErrorException(err.Error()))
	}

	return users, int(total)
}

// func filter for pagination
func (user *UserRepository) queryFilter(query *gorm.DB, httpContext *gin.Context) *gorm.DB {
	if search := httpContext.Query("search"); search != "" {
		fmt.Println("search", search)
		query = query.Where("name LIKE ?", search+"%")
	}

	return query
}

// func query sort for pagination
func (user *UserRepository) querySort(query *gorm.DB, httpContext *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "email", "created_at", "updated_at"}

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

func (user *UserRepository) FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) *domain.UserEntity {
	user.model = user.model.WithContext(httpContext)
	userEntity := &domain.UserEntity{}
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
		panic(*exception.ServerErrorException(err.Error()))
	}

	return userEntity
}

func (user *UserRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	user.model = user.model.WithContext(httpContext)
	err := user.model.Delete(&domain.UserEntity{}, id).Error

	if err != nil {
		log.Println("Error delete user", err)
		panic(*exception.ServerErrorException(err.Error()))
	}
}

func (user *UserRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.UserEntity) {
	user.model = user.model.WithContext(httpContext)
	userEntity := &domain.UserEntity{}

	err := user.model.Transaction(func(tx *gorm.DB) error {
		// Find user within transaction
		if err := tx.First(&userEntity, id).Error; err != nil {
			return err
		}

		// Update roles within transaction
		if err := tx.Model(&userEntity).Association("Roles").Replace(payload.Roles); err != nil {
			return err
		}

		// Update user data within transaction
		if err := tx.Where("id = ?", id).Updates(&payload).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println("Error update user", err)
		panic(*exception.ServerErrorException(err.Error()))
	}
}

func (user *UserRepository) Create(httpContext *gin.Context, payload *domain.UserEntity) {
	user.model = user.model.WithContext(httpContext)
	err := user.model.Create(&payload).Error

	if err != nil {
		log.Println("Error create user", err)
		panic(*exception.ServerErrorException(err.Error()))
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

func (user *UserRepository) FindRoleByIds(httpContext *gin.Context, ids []uuid.UUID) []roleDomain.RoleEntity {
	user.roleModel = user.roleModel.WithContext(httpContext)
	var roleEntities []roleDomain.RoleEntity
	err := user.roleModel.Where("id IN ?", ids).Find(&roleEntities).Error

	if err != nil {
		log.Println("Error finding roles", err)
		panic(*exception.BussinessException(err.Error()))
	}

	if len(roleEntities) == 0 {
		panic(*exception.NotFoundException("Roles not found"))
	}

	return roleEntities
}

// Delete user's roles using association table without deleting the user
func (user *UserRepository) DeleteUserRoles(httpContext *gin.Context, id uuid.UUID) {
	user.model = user.model.WithContext(httpContext)
	err := user.model.Association("Roles").Clear()

	if err != nil {
		log.Println("Error delete user roles", err)
		panic(*exception.ServerErrorException(err.Error()))
	}
}
