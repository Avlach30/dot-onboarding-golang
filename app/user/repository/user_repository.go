package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	permissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	roleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	rolePermissionDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role_permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	userModel           *gorm.DB
	permissionModel     *gorm.DB
	roleModel           *gorm.DB
	rolePermissionModel *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{
		userModel:           db.Model(&userDomain.UserEntity{}),
		permissionModel:     db.Model(&permissionDomain.PermissionEntity{}),
		roleModel:           db.Model(&roleDomain.RoleEntity{}),
		rolePermissionModel: db.Model(&rolePermissionDomain.RolePermissionEntity{}),
	}
}

// Pagination get user data with pagination
func (user *UserRepository) Pagination(ctx *gin.Context) ([]domain.UserEntity, int) {
	user.userModel = user.userModel.WithContext(ctx)
	var users []domain.UserEntity
	var total int64

	// Query filter
	user.queryFilter(ctx)
	// Query sort
	user.querySort(ctx)

	user.userModel.Session(&gorm.Session{}).
		Scopes(utils.Paginate(ctx)).
		Find(&users).
		Count(&total)

	return users, int(total)
}

// func filter for pagination
func (user *UserRepository) queryFilter(ctx *gin.Context) *gorm.DB {
	if search := ctx.Query("search"); search != "" {
		user.userModel = user.userModel.
			Where("name LIKE ?", "%"+search+"%")
	}

	return user.userModel
}

// func query sort for pagination
func (user *UserRepository) querySort(ctx *gin.Context) *gorm.DB {
	sortableColumns := []string{"name", "email", "created_at", "updated_at"}

	if sort := ctx.Query("sort_by"); sort != "" {
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		user.userModel = user.userModel.
			Order(sort + " " + ctx.Query("order"))

		// Handle order by asc or desc
		if order := ctx.Query("order"); order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}

			user.userModel = user.userModel.
				Order(sort + " " + order)
		}
	}

	return user.userModel
}

func (user *UserRepository) FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*domain.UserEntity, error) {
	user.userModel = user.userModel.WithContext(ctx)
	userEntity := &domain.UserEntity{}
	if trashed {
		user.userModel = user.userModel.Unscoped()
	}

	err := user.userModel.Where("id = ?", id).First(&userEntity).Error

	return userEntity, err
}

func (user *UserRepository) Delete(ctx *gin.Context, id uuid.UUID) {
	user.userModel = user.userModel.WithContext(ctx)
	user.userModel.Delete(&domain.UserEntity{}, id)
}

func (user *UserRepository) ForceDelete(ctx *gin.Context, id uuid.UUID) {
	user.userModel = user.userModel.WithContext(ctx)
	userEntity := &domain.UserEntity{}
	user.userModel.Unscoped().Delete(&userEntity, id)
}

func (user *UserRepository) Update(ctx *gin.Context, id uuid.UUID, payload *domain.UserEntity) error {
	user.userModel = user.userModel.WithContext(ctx)
	err := user.userModel.Where("id = ?", id).Updates(&payload).Error

	return err
}

func (user *UserRepository) Create(ctx *gin.Context, payload *domain.UserEntity) error {
	user.userModel = user.userModel.WithContext(ctx)
	err := user.userModel.Create(&payload).Error
	return err
}

func (user *UserRepository) IsEmailExist(ctx *gin.Context, email string) bool {
	user.userModel = user.userModel.WithContext(ctx)
	var count int64
	user.userModel.
		Where("email = ?", email).
		Count(&count)
	return count > 0
}

func (user *UserRepository) IsEmailExistExceptUserId(ctx *gin.Context, email string, id uuid.UUID) bool {
	user.userModel = user.userModel.WithContext(ctx)
	var count int64
	user.userModel.
		Where("email = ? AND id != ?", email, id).
		Count(&count)

	return count > 0
}
