package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type PermissionRepository interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.PermissionEntity, int)
	Create(httpContext *gin.Context, payload *entities.PermissionEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.PermissionEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.PermissionEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsKeyExist(httpContext *gin.Context, key string) bool
	IsKeyExistExceptPermissionId(httpContext *gin.Context, key string, id uuid.UUID) bool
}
