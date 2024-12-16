package usecase

import (
	"context"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

// Create implements domain.UserUsecase.
func (userUsecase *UserUsecase) Create(context *context.Context, payload *domain.UserEntity) error {
	isUserExist := userUsecase.userRepo.IsEmailExist(context, payload.Email)

	if isUserExist {
		panic(*exception.BussinessException("Email already exist"))
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(*exception.ServerErrorException("Failed to hash password"))
	}

	payload.Password = string(hashedPassword)

	return userUsecase.userRepo.Create(context, payload)
}

// Delete implements domain.UserUsecase.
func (userUsecase *UserUsecase) Delete(context *context.Context, id uuid.UUID) {
	userUsecase.userRepo.Delete(context, id)
}

// FindById implements domain.UserUsecase.
func (userUsecase *UserUsecase) FindById(context *context.Context, id uuid.UUID, trashed bool) (*domain.UserEntity, error) {
	user, err := userUsecase.userRepo.FindById(context, id, trashed)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("User not found"))
	}

	return user, err
}

// ForceDelete implements domain.UserUsecase.
func (userUsecase *UserUsecase) ForceDelete(context *context.Context, id uuid.UUID) {
	userUsecase.userRepo.ForceDelete(context, id)
}

// Update implements domain.UserUsecase.
func (userUsecase *UserUsecase) Update(context *context.Context, id uuid.UUID, payload *domain.UserEntity) {
	if userUsecase.userRepo.IsEmailExistExceptUserId(context, payload.Email, id) {
		panic(*exception.BussinessException("Email already exist"))
	}

	err := userUsecase.userRepo.Update(context, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update user"))
	}
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
