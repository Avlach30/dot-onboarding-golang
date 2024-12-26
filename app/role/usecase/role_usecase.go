package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
)

type RoleUsecase struct {
	roleRepo domain.RoleRepository
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}

func (roleUsecase *RoleUsecase) Pagination(httpContext *gin.Context) ([]domain.RoleEntity, int) {
	return roleUsecase.roleRepo.Pagination(httpContext)
}

// Create implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Create(httpContext *gin.Context, payload *domain.RoleEntity) {
	isKeyExist := roleUsecase.roleRepo.IsKeyExist(httpContext, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	roleUsecase.roleRepo.Create(httpContext, payload)
}

// Delete implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	roleUsecase.roleRepo.Delete(httpContext, id)
}

// FindOneById implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID) *domain.RoleEntity {
	return roleUsecase.roleRepo.FindOneById(httpContext, id, false)
}

// Update implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.RoleEntity) {
	if roleUsecase.roleRepo.IsKeyExistExceptRoleId(httpContext, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	roleUsecase.roleRepo.Update(httpContext, id, payload)
}
