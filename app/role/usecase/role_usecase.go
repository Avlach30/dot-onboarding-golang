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

func (roleUsecase *RoleUsecase) Pagination(httpContext *gin.Context) ([]domain.RoleEntity, int) {
	return roleUsecase.roleRepo.Pagination(httpContext)
}

// Create implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Create(httpContext *gin.Context, payload *domain.RoleEntity) error {
	isKeyExist := roleUsecase.roleRepo.IsKeyExist(httpContext, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return roleUsecase.roleRepo.Create(httpContext, payload)
}

// Delete implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	roleUsecase.roleRepo.Delete(httpContext, id)
}

// FindById implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindById(httpContext *gin.Context, id uuid.UUID) (*domain.RoleEntity, error) {
	role, err := roleUsecase.roleRepo.FindById(httpContext, id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Role not found"))
	}

	return role, err
}

// FindByKey implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindByKey(httpContext *gin.Context, key string) (*domain.RoleEntity, error) {
	return roleUsecase.roleRepo.FindByKey(httpContext, key, false)
}

// Update implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.RoleEntity) {
	if roleUsecase.roleRepo.IsKeyExistExceptRoleId(httpContext, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := roleUsecase.roleRepo.Update(httpContext, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update role"))
	}
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}
