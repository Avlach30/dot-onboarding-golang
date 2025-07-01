package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type PermissionUsecase interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.PermissionEntity, int)
	Create(httpContext *gin.Context, schema *entities.PermissionEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID) *entities.PermissionEntity
	Update(httpContext *gin.Context, id uuid.UUID, schema *entities.PermissionEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
