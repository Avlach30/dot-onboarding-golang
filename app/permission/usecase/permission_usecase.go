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

func (permissionUsecase *PermissionUsecase) Pagination(ctx *gin.Context) ([]domain.PermissionEntity, int) {
	return permissionUsecase.permissionRepo.Pagination(ctx)
}

// Create implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Create(ctx *gin.Context, payload *domain.PermissionEntity) error {
	isKeyExist := permissionUsecase.permissionRepo.IsKeyExist(ctx, payload.Key)

	if isKeyExist {
		panic(*exception.BussinessException("Key already exist"))
	}

	return permissionUsecase.permissionRepo.Create(ctx, payload)
}

// Delete implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Delete(ctx *gin.Context, id uuid.UUID) {
	permissionUsecase.permissionRepo.Delete(ctx, id)
}

// FindById implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindById(ctx *gin.Context, id uuid.UUID) (*domain.PermissionEntity, error) {
	permission, err := permissionUsecase.permissionRepo.FindById(ctx, id, false)

	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Permission not found"))
	}

	return permission, err
}

// FindByKey implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) FindByKey(ctx *gin.Context, key string) (*domain.PermissionEntity, error) {
	return permissionUsecase.permissionRepo.FindByKey(ctx, key, false)
}

// Update implements domain.PermissionUsecase.
func (permissionUsecase *PermissionUsecase) Update(ctx *gin.Context, id uuid.UUID, payload *domain.PermissionEntity) {
	if permissionUsecase.permissionRepo.IsKeyExistExceptPermissionId(ctx, payload.Key, id) {
		panic(*exception.BussinessException("Key already exist"))
	}

	err := permissionUsecase.permissionRepo.Update(ctx, id, payload)
	if err != nil {
		panic(*exception.ServerErrorException("Failed to update permission"))
	}
}

func NewPermissionUsecase(permissionRepo domain.PermissionRepository) domain.PermissionUsecase {
	return &PermissionUsecase{
		permissionRepo: permissionRepo,
	}
}
