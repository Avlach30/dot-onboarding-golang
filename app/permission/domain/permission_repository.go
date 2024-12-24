package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	Pagination(httpContext *gin.Context) ([]PermissionEntity, int)
	Create(httpContext *gin.Context, payload *PermissionEntity) error
	FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) (*PermissionEntity, error)
	FindByKey(httpContext *gin.Context, key string, trashed bool) (*PermissionEntity, error)
	FindByNameAndKey(httpContext *gin.Context, name string, key string) (*PermissionEntity, error)
	Update(httpContext *gin.Context, id uuid.UUID, payload *PermissionEntity) error
	Delete(httpContext *gin.Context, id uuid.UUID)
	ForceDelete(httpContext *gin.Context, id uuid.UUID)
	IsKeyExist(httpContext *gin.Context, key string) bool
	IsKeyExistExceptPermissionId(httpContext *gin.Context, key string, id uuid.UUID) bool
}
