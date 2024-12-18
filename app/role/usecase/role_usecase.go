package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/gorm"
)

type RoleUsecase struct {
	roleRepo domain.RoleRepository
}

// Create implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Create(ctx *gin.Context, payload *domain.RoleEntity) error {
	isKeyExist := roleUsecase.roleRepo.IsKeyExist(ctx, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return roleUsecase.roleRepo.Create(ctx, payload)
}

// Delete implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Delete(ctx *gin.Context, id uuid.UUID) {
	roleUsecase.roleRepo.Delete(ctx, id)
}

// FindById implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindById(ctx *gin.Context, id uuid.UUID) (*domain.RoleEntity, error) {
	role, err := roleUsecase.roleRepo.FindById(ctx, id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Role not found"))
	}

	return role, err
}

// FindByKey implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindByKey(ctx *gin.Context, key string) (*domain.RoleEntity, error) {
	return roleUsecase.roleRepo.FindByKey(ctx, key, false)
}

// Update implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Update(ctx *gin.Context, id uuid.UUID, payload *domain.RoleEntity) {
	if roleUsecase.roleRepo.IsKeyExistExceptRoleId(ctx, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := roleUsecase.roleRepo.Update(ctx, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update role"))
	}
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}
