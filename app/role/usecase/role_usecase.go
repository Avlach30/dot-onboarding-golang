package usecase

import (
	"context"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/gorm"
)

type RoleUsecase struct {
	roleRepo domain.RoleRepository
}

// Create implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Create(context *context.Context, payload *domain.RoleEntity) error {
	isKeyExist := roleUsecase.roleRepo.IsKeyExist(context, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return roleUsecase.roleRepo.Create(context, payload)
}

// Delete implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Delete(context *context.Context, id uuid.UUID) {
	roleUsecase.roleRepo.Delete(context, id)
}

// FindById implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindById(context *context.Context, id uuid.UUID) (*domain.RoleEntity, error) {
	role, err := roleUsecase.roleRepo.FindById(context, id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Role not found"))
	}

	return role, err
}

// FindByKey implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindByKey(context *context.Context, key string) (*domain.RoleEntity, error) {
	return roleUsecase.roleRepo.FindByKey(context, key, false)
}

// Update implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Update(context *context.Context, id uuid.UUID, payload *domain.RoleEntity) {
	if roleUsecase.roleRepo.IsKeyExistExceptRoleId(context, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := roleUsecase.roleRepo.Update(context, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update role"))
	}
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}
