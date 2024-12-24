package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleRepository interface {
	Pagination(httpContext *gin.Context) ([]RoleEntity, int)
	Create(httpContext *gin.Context, payload *RoleEntity) error
	FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) (*RoleEntity, error)
	FindByKey(httpContext *gin.Context, key string, trashed bool) (*RoleEntity, error)
	FindByNameAndKey(httpContext *gin.Context, name string, key string) (*RoleEntity, error)
	Update(httpContext *gin.Context, id uuid.UUID, payload *RoleEntity) error
	Delete(httpContext *gin.Context, id uuid.UUID)
	ForceDelete(httpContext *gin.Context, id uuid.UUID)
	IsKeyExist(httpContext *gin.Context, key string) bool
	IsKeyExistExceptRoleId(httpContext *gin.Context, key string, id uuid.UUID) bool
}
