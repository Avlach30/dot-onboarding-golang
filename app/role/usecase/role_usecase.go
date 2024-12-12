package usecase

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
)

type RoleUsecase struct {
	roleRepo domain.RoleRepository
}

// Create implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Create(schema *domain.RoleEntity) error {
	return roleUsecase.roleRepo.Create(schema)
}

// Delete implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Delete(id uuid.UUID) {
	roleUsecase.roleRepo.Delete(id)
}

// FindById implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindById(id uuid.UUID) (*domain.RoleEntity, error) {
	return roleUsecase.roleRepo.FindById(id, false)
}

// FindByKey implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindByKey(key string) (*domain.RoleEntity, error) {
	return roleUsecase.roleRepo.FindByKey(key, false)
}

// Update implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Update(id uuid.UUID, schema *domain.RoleEntity) {
	roleUsecase.roleRepo.Update(id, schema)
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}
