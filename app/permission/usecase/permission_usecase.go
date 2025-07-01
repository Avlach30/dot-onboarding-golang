package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionUsecase struct {
	permissionRepo domain.PermissionRepository
}

func NewPermissionUsecase(permissionRepo domain.PermissionRepository) domain.PermissionUsecase {
	return &PermissionUsecase{
		permissionRepo: permissionRepo,
	}
}

func (permissionUsecase *PermissionUsecase) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.PermissionEntity, int) {
	return permissionUsecase.permissionRepo.Pagination(httpContext, queryDto)
}

// Create implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Create(httpContext *gin.Context, payload *entities.PermissionEntity) {
	isKeyExist := permissionUsecase.permissionRepo.IsKeyExist(httpContext, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	permissionUsecase.permissionRepo.Create(httpContext, payload)
}

// Delete implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	permissionUsecase.permissionRepo.Delete(httpContext, id)
}

// FindOneById implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID) *entities.PermissionEntity {
	permission := permissionUsecase.permissionRepo.FindOneById(httpContext, id, false)

	return permission
}

// Update implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.PermissionEntity) {
	if permissionUsecase.permissionRepo.IsKeyExistExceptPermissionId(httpContext, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	permissionUsecase.permissionRepo.Update(httpContext, id, payload)
}
