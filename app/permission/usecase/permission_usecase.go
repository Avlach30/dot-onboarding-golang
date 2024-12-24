package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionUsecase struct {
	permissionRepo domain.PermissionRepository
}

func (permissionUsecase *PermissionUsecase) Pagination(httpContext *gin.Context) ([]domain.PermissionEntity, int) {
	return permissionUsecase.permissionRepo.Pagination(httpContext)
}

// Create implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Create(httpContext *gin.Context, payload *domain.PermissionEntity) error {
	isKeyExist := permissionUsecase.permissionRepo.IsKeyExist(httpContext, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return permissionUsecase.permissionRepo.Create(httpContext, payload)
}

// Delete implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	permissionUsecase.permissionRepo.Delete(httpContext, id)
}

// FindById implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindById(httpContext *gin.Context, id uuid.UUID) (*domain.PermissionEntity, error) {
	permission, err := permissionUsecase.permissionRepo.FindById(httpContext, id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Permission not found"))
	}

	return permission, err
}

// FindByKey implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindByKey(httpContext *gin.Context, key string) (*domain.PermissionEntity, error) {
	return permissionUsecase.permissionRepo.FindByKey(httpContext, key, false)
}

// Update implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *domain.PermissionEntity) {
	if permissionUsecase.permissionRepo.IsKeyExistExceptPermissionId(httpContext, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := permissionUsecase.permissionRepo.Update(httpContext, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update permission"))
	}
}

func NewPermissionUsecase(permissionRepo domain.PermissionRepository) domain.PermissionUsecase {
	return &PermissionUsecase{
		permissionRepo: permissionRepo,
	}
}
