package usecase

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/gorm"
)

type RoleUsecase struct {
	roleRepo domain.RoleRepository
}

// Create implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Create(payload *domain.RoleEntity) error {
	isKeyExist := roleUsecase.roleRepo.IsKeyExist(payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return roleUsecase.roleRepo.Create(payload)
}

// Delete implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Delete(id uuid.UUID) {
	roleUsecase.roleRepo.Delete(id)
}

// FindById implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindById(id uuid.UUID) (*domain.RoleEntity, error) {
	role, err := roleUsecase.roleRepo.FindById(id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Role not found"))
	}

	return role, err
}

// FindByKey implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindByKey(key string) (*domain.RoleEntity, error) {
	return roleUsecase.roleRepo.FindByKey(key, false)
}

// Update implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Update(id uuid.UUID, payload *domain.RoleEntity) {
	if roleUsecase.roleRepo.IsKeyExistExceptRoleId(payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := roleUsecase.roleRepo.Update(id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update role"))
	}
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}
