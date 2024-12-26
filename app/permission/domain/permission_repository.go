package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionRepository interface {
	Pagination(httpContext *gin.Context) ([]PermissionEntity, int)
	Create(httpContext *gin.Context, payload *PermissionEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *PermissionEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *PermissionEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsKeyExist(httpContext *gin.Context, key string) bool
	IsKeyExistExceptPermissionId(httpContext *gin.Context, key string, id uuid.UUID) bool
}
