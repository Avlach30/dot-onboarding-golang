package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type RoleUsecase struct {
	roleRepo domain.RoleRepository
}

func NewRoleUsecase(roleRepo domain.RoleRepository) domain.RoleUsecase {
	return &RoleUsecase{
		roleRepo: roleRepo,
	}
}

func (roleUsecase *RoleUsecase) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.RoleEntity, int) {
	return roleUsecase.roleRepo.Pagination(httpContext, queryDto)
}

// Create implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Create(httpContext *gin.Context, payload *entities.RoleEntity) {
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
func (roleUsecase *RoleUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID) *entities.RoleEntity {
	return roleUsecase.roleRepo.FindOneById(httpContext, id, false)
}

// Update implements domain.RoleUsecase.
func (roleUsecase *RoleUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.RoleEntity) {
	if roleUsecase.roleRepo.IsKeyExistExceptRoleId(httpContext, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	roleUsecase.roleRepo.Update(httpContext, id, payload)
}
