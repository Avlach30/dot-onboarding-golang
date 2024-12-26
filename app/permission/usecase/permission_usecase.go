package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"

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

func (permissionUsecase *PermissionUsecase) Pagination(httpContext *gin.Context) ([]domain.PermissionEntity, int) {
	return permissionUsecase.permissionRepo.Pagination(httpContext)
}

// Create implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Create(httpContext *gin.Context, payload *domain.PermissionEntity) {
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
func (permissionUsecase *PermissionUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID) *domain.PermissionEntity {
	permission := permissionUsecase.permissionRepo.FindOneById(httpContext, id, false)

	return permission
}

// Update implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.PermissionEntity) {
	if permissionUsecase.permissionRepo.IsKeyExistExceptPermissionId(httpContext, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	permissionUsecase.permissionRepo.Update(httpContext, id, payload)
}
