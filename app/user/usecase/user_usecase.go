package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

// Create implements domain.UserUsecase.
func (userUsecase *UserUsecase) Create(payload *domain.UserEntity) error {
	return userUsecase.userRepo.Create(payload)
}

// Delete implements domain.UserUsecase.
func (userUsecase *UserUsecase) Delete(id uuid.UUID) {
	userUsecase.userRepo.Delete(id)
}

// FindById implements domain.UserUsecase.
func (userUsecase *UserUsecase) FindById(id uuid.UUID, trashed bool) (*domain.UserEntity, error) {
	user, err := userUsecase.userRepo.FindById(id, trashed)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("User not found"))
	}

	return user, err
}

// ForceDelete implements domain.UserUsecase.
func (userUsecase *UserUsecase) ForceDelete(id uuid.UUID) {
	userUsecase.userRepo.ForceDelete(id)
}

// Update implements domain.UserUsecase.
func (userUsecase *UserUsecase) Update(id uuid.UUID, payload *domain.UserEntity) {
	userUsecase.userRepo.Update(id, payload)
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}
