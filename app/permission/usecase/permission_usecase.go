package usecase

import (
	"context"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/gorm"

	"github.com/google/uuid"
)

type PermissionUsecase struct {
	permissionRepo domain.PermissionRepository
}

// Create implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Create(context *context.Context, payload *domain.PermissionEntity) error {
	isKeyExist := permissionUsecase.permissionRepo.IsKeyExist(context, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return permissionUsecase.permissionRepo.Create(context, payload)
}

// Delete implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Delete(context *context.Context, id uuid.UUID) {
	permissionUsecase.permissionRepo.Delete(context, id)
}

// FindById implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindById(context *context.Context, id uuid.UUID) (*domain.PermissionEntity, error) {
	permission, err := permissionUsecase.permissionRepo.FindById(context, id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Permission not found"))
	}

	return permission, err
}

// FindByKey implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindByKey(context *context.Context, key string) (*domain.PermissionEntity, error) {
	return permissionUsecase.permissionRepo.FindByKey(context, key, false)
}

// Update implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Update(context *context.Context, id uuid.UUID, payload *domain.PermissionEntity) {
	if permissionUsecase.permissionRepo.IsKeyExistExceptPermissionId(context, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := permissionUsecase.permissionRepo.Update(context, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update permission"))
	}
}

func NewPermissionUsecase(permissionRepo domain.PermissionRepository) domain.PermissionUsecase {
	return &PermissionUsecase{
		permissionRepo: permissionRepo,
	}
}
