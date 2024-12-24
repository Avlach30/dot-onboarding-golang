package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleUsecase interface {
	Pagination(httpContext *gin.Context) ([]RoleEntity, int)
	Create(httpContext *gin.Context, schema *RoleEntity) error
	FindById(httpContext *gin.Context, id uuid.UUID) (*RoleEntity, error)
	FindByKey(httpContext *gin.Context, key string) (*RoleEntity, error)
	Update(httpContext *gin.Context, id uuid.UUID, dto *RoleEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
