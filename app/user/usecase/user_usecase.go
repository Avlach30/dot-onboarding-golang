package usecase

import (
	"log"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"golang.org/x/crypto/bcrypt"

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
func (userUsecase *UserUsecase) Pagination(httpContext *gin.Context) ([]domain.UserEntity, int) {
	return userUsecase.userRepo.Pagination(httpContext)
}

// Create implements domain.UserUsecase.
func (userUsecase *UserUsecase) Create(httpContext *gin.Context, payload *domain.UserEntity, roleIds []uuid.UUID) {
	isUserExist := userUsecase.userRepo.IsEmailExist(httpContext, payload.Email)
	if isUserExist {
		panic(*exception.BussinessException("Email already exist"))
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password", err)
		panic(*exception.ServerErrorException(err))
	}

	payload.Password = string(hashedPassword)

	// Get role by role ids
	roles := userUsecase.userRepo.FindRoleByIds(httpContext, roleIds)
	// Validate roles length is same with role ids
	if len(roles) != len(roleIds) {
		panic(*exception.BussinessException("Role not found"))
	}

	// Assign payload
	payload.Roles = roles

	userUsecase.userRepo.Create(httpContext, payload)
}

// Update implements domain.UserUsecase.
func (userUsecase *UserUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *dto.UserUpdateRequest) {
	if userUsecase.userRepo.IsEmailExistExceptUserId(httpContext, payload.Email, id) {
		panic(*exception.BussinessException("Email already exist"))
	}

	// Get role by role ids
	roles := userUsecase.userRepo.FindRoleByIds(httpContext, payload.RoleIds)
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

	userUsecase.userRepo.Update(httpContext, id, &payloadEntity)

}

// Delete implements domain.UserUsecase.
func (userUsecase *UserUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	userUsecase.userRepo.Delete(httpContext, id)
}

// FindById implements domain.UserUsecase.
func (userUsecase *UserUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *domain.UserEntity {
	return userUsecase.userRepo.FindOneById(httpContext, id, trashed)
}
