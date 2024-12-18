package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

// Pagination implements domain.UserUsecase.
func (userUsecase *UserUsecase) Pagination(ctx *gin.Context) ([]domain.User, int) {
	return userUsecase.userRepo.Pagination(ctx)
}

// Create implements domain.UserUsecase.
func (userUsecase *UserUsecase) Create(ctx *gin.Context, payload *domain.User) error {
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

	return userUsecase.userRepo.Create(ctx, payload)
}

// Delete implements domain.UserUsecase.
func (userUsecase *UserUsecase) Delete(ctx *gin.Context, id uuid.UUID) {
	userUsecase.userRepo.Delete(ctx, id)
}

// FindById implements domain.UserUsecase.
func (userUsecase *UserUsecase) FindById(ctx *gin.Context, id uuid.UUID, trashed bool) (*domain.User, error) {
	user, err := userUsecase.userRepo.FindById(ctx, id, trashed)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("User not found"))
	}

	return user, err
}

// ForceDelete implements domain.UserUsecase.
func (userUsecase *UserUsecase) ForceDelete(ctx *gin.Context, id uuid.UUID) {
	userUsecase.userRepo.ForceDelete(ctx, id)
}

// Update implements domain.UserUsecase.
func (userUsecase *UserUsecase) Update(ctx *gin.Context, id uuid.UUID, payload *domain.User) {
	if userUsecase.userRepo.IsEmailExistExceptUserId(ctx, payload.Email, id) {
		panic(*exception.BussinessException("Email already exist"))
	}

	err := userUsecase.userRepo.Update(ctx, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update user"))
	}
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
