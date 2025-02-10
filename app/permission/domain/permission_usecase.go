package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/entities"
)

type PermissionUsecase interface {
	Pagination(httpContext *gin.Context) ([]entities.PermissionEntity, int)
	Create(httpContext *gin.Context, schema *entities.PermissionEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID) *entities.PermissionEntity
	Update(httpContext *gin.Context, id uuid.UUID, schema *entities.PermissionEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
