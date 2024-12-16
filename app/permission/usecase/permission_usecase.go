package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type PermissionUsecase struct {
	permissionRepo domain.PermissionRepository
}

// Create implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Create(payload *domain.PermissionEntity) error {
	isKeyExist := permissionUsecase.permissionRepo.IsKeyExist(payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return permissionUsecase.permissionRepo.Create(payload)
}

// Delete implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Delete(id uuid.UUID) {
	permissionUsecase.permissionRepo.Delete(id)
}

// FindById implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindById(id uuid.UUID) (*domain.PermissionEntity, error) {
	permission, err := permissionUsecase.permissionRepo.FindById(id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Permission not found"))
	}

	return permission, err
}

// FindByKey implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindByKey(key string) (*domain.PermissionEntity, error) {
	return permissionUsecase.permissionRepo.FindByKey(key, false)
}

// Update implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Update(id uuid.UUID, payload *domain.PermissionEntity) {
	if permissionUsecase.permissionRepo.IsKeyExistExceptPermissionId(payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := permissionUsecase.permissionRepo.Update(id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update permission"))
	}
}

func NewPermissionUsecase(permissionRepo domain.PermissionRepository) domain.PermissionUsecase {
	return &PermissionUsecase{
		permissionRepo: permissionRepo,
	}
}
