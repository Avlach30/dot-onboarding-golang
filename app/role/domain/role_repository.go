package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleRepository interface {
	Pagination(httpContext *gin.Context) ([]RoleEntity, int)
	Create(httpContext *gin.Context, payload *RoleEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *RoleEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *RoleEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsKeyExist(httpContext *gin.Context, key string) bool
	IsKeyExistExceptRoleId(httpContext *gin.Context, key string, id uuid.UUID) bool
}
