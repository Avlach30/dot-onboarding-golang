package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	userModel *gorm.DB
	roleModel *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		userModel: db.Model(&domain.UserEntity{}),
		roleModel: db.Model(&roleDomain.RoleEntity{}),
	}
}

// Pagination get user data with pagination
func (user *UserRepository) Pagination(httpContext *gin.Context) ([]domain.UserEntity, int) {
	user.userModel = user.userModel.WithContext(httpContext)
	var users []domain.UserEntity
	var total int64

	// Query filter
	user.queryFilter(httpContext)
	// Query sort
	user.querySort(httpContext)

	user.userModel.Session(&gorm.Session{}).
		Scopes(utils.Paginate(httpContext)).
		Find(&users).
		Count(&total)

	return users, int(total)
}

// func filter for pagination
func (user *UserRepository) queryFilter(httpContext *gin.Context) *gorm.DB {
	if search := httpContext.Query("search"); search != "" {
		user.userModel = user.userModel.
			Where("name LIKE ?", "%"+search+"%")
	}

	return user.userModel
}

// func query sort for pagination
func (user *UserRepository) querySort(httpContext *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "email", "created_at", "updated_at"}

	if sort := httpContext.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		user.userModel = user.userModel.
			Order(sort + " " + httpContext.Query("order"))

		// Handle order by asc or desc
		if order := httpContext.Query("order"); order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}

			user.userModel = user.userModel.
				Order(sort + " " + order)
		}
	}

	return user.userModel
}

func (user *UserRepository) FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) (*domain.UserEntity, error) {
	user.userModel = user.userModel.WithContext(httpContext)
	userEntity := &domain.UserEntity{}
	if trashed {
		user.userModel = user.userModel.Unscoped()
	}

	err := user.userModel.
		Preload("Roles").
		Preload("Roles.Permissions").
		First(&userEntity, id).
		Error

	return userEntity, err
}

func (user *UserRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	user.userModel = user.userModel.WithContext(httpContext)
	user.userModel.Delete(&domain.UserEntity{}, id)
}

func (user *UserRepository) ForceDelete(httpContext *gin.Context, id uuid.UUID) {
	user.userModel = user.userModel.WithContext(httpContext)
	userEntity := &domain.UserEntity{}
	user.userModel.Unscoped().Delete(&userEntity, id)
}

func (user *UserRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.UserEntity) {
	user.userModel = user.userModel.WithContext(httpContext)
	userEntity := &domain.UserEntity{}

	err := user.userModel.Transaction(func(tx *gorm.DB) error {
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
		panic(*exception.ServerErrorException("Failed to update user"))
	}
}

func (user *UserRepository) Create(httpContext *gin.Context, payload *domain.UserEntity) {
	user.userModel = user.userModel.WithContext(httpContext)
	err := user.userModel.Create(&payload).Error

	if err != nil {
		panic(*exception.ServerErrorException("Failed to create user"))
	}
}

func (user *UserRepository) IsEmailExist(httpContext *gin.Context, email string) bool {
	user.userModel = user.userModel.WithContext(httpContext)
	var count int64
	user.userModel.
		Where("email = ?", email).
		Count(&count)
	return count > 0
}

func (user *UserRepository) IsEmailExistExceptUserId(httpContext *gin.Context, email string, id uuid.UUID) bool {
	user.userModel = user.userModel.WithContext(httpContext)
	var count int64
	user.userModel.
		Where("email = ? AND id != ?", email, id).
		Count(&count)

	return count > 0
}

func (user *UserRepository) FindRoleByIds(httpContext *gin.Context, ids []uuid.UUID) []roleDomain.RoleEntity {
	user.roleModel = user.roleModel.WithContext(httpContext)
	var roleEntities []roleDomain.RoleEntity
	err := user.roleModel.Where("id IN ?", ids).Find(&roleEntities).Error

	if err != nil {
		panic(*exception.BussinessException("Error finding roles"))
	}

	if len(roleEntities) == 0 {
		panic(*exception.NotFoundException("Roles not found"))
	}

	return roleEntities
}

// Delete user's roles using association table without deleting the user
func (user *UserRepository) DeleteUserRoles(httpContext *gin.Context, id uuid.UUID) {
	user.userModel = user.userModel.WithContext(httpContext)
	user.userModel.Association("Roles").Clear()
}
