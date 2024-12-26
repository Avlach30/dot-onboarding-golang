package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionUsecase interface {
	Pagination(httpContext *gin.Context) ([]PermissionEntity, int)
	Create(httpContext *gin.Context, schema *PermissionEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID) *PermissionEntity
	Update(httpContext *gin.Context, id uuid.UUID, schema *PermissionEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
