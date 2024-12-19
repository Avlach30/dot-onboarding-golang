package usecase

import (
	"fmt"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

// Pagination implements domain.UserUsecase.
func (userUsecase *UserUsecase) Pagination(ctx *gin.Context) ([]domain.UserEntity, int) {
	return userUsecase.userRepo.Pagination(ctx)
}

// Create implements domain.UserUsecase.
func (userUsecase *UserUsecase) Create(ctx *gin.Context, payload *dto.UserCreateRequest) {
	isUserExist := userUsecase.userRepo.IsEmailExist(ctx, payload.Email)
	if isUserExist {
		panic(*exception.BussinessException("Email already exist"))
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to hash password"))
	}

	payload.Password = string(hashedPassword)

	// Get role by role ids
	roles := userUsecase.userRepo.FindRoleByIds(ctx, payload.RoleIds)
	// Validate roles length is same with role ids
	if len(roles) != len(payload.RoleIds) {
		panic(*exception.BussinessException("Role not found"))
	}

	// Extract payload to entity
	payloadEntity := domain.UserEntity{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: payload.Password,
		Roles:    roles,
	}

	userUsecase.userRepo.Create(ctx, &payloadEntity)
}

// Update implements domain.UserUsecase.
func (userUsecase *UserUsecase) Update(ctx *gin.Context, id uuid.UUID, payload *dto.UserUpdateRequest) {
	if userUsecase.userRepo.IsEmailExistExceptUserId(ctx, payload.Email, id) {
		panic(*exception.BussinessException("Email already exist"))
	}

	// Get role by role ids
	roles := userUsecase.userRepo.FindRoleByIds(ctx, payload.RoleIds)
	// Validate roles length is same with role ids
	if len(roles) != len(payload.RoleIds) {
		panic(*exception.BussinessException("Role not found"))
	}

	// Extract payload to entity
	payloadEntity := domain.UserEntity{
		Name:  payload.Name,
		Email: payload.Email,
		Roles: roles,
	}

	userUsecase.userRepo.Update(ctx, id, &payloadEntity)

}

// Delete implements domain.UserUsecase.
func (userUsecase *UserUsecase) Delete(ctx *gin.Context, id uuid.UUID) {
	userUsecase.userRepo.Delete(ctx, id)
}

// FindById implements domain.UserUsecase.
func (userUsecase *UserUsecase) FindById(ctx *gin.Context, id uuid.UUID, trashed bool) *domain.UserEntity {
	user, err := userUsecase.userRepo.FindById(ctx, id, trashed)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("User not found"))
	} else if err != nil {
		fmt.Println(err)
		panic(*exception.ServerErrorException("Failed to find user"))
	}

	return user
}

// ForceDelete implements domain.UserUsecase.
func (userUsecase *UserUsecase) ForceDelete(ctx *gin.Context, id uuid.UUID) {
	userUsecase.userRepo.ForceDelete(ctx, id)
}
