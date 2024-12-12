package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"

	"github.com/google/uuid"
)

type PermissionUsecase struct {
	permissionRepo domain.PermissionRepository
}

// Create implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Create(schema *domain.PermissionEntity) error {
	return permissionUsecase.permissionRepo.Create(schema)
}

// Delete implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Delete(id uuid.UUID) {
	permissionUsecase.permissionRepo.Delete(id)
}

// FindById implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindById(id uuid.UUID) (*domain.PermissionEntity, error) {
	return permissionUsecase.permissionRepo.FindById(id, false)
}

// FindByKey implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindByKey(key string) (*domain.PermissionEntity, error) {
	return permissionUsecase.permissionRepo.FindByKey(key, false)
}

// Update implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Update(id uuid.UUID, schema *domain.PermissionEntity) {
	permissionUsecase.permissionRepo.Update(id, schema)
}

func NewPermissionUsecase(permissionRepo domain.PermissionRepository) domain.PermissionUsecase {
	return &PermissionUsecase{
		permissionRepo: permissionRepo,
	}
}
