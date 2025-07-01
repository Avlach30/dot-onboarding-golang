package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type RoleUsecase interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.RoleEntity, int)
	Create(httpContext *gin.Context, schema *entities.RoleEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID) *entities.RoleEntity
	Update(httpContext *gin.Context, id uuid.UUID, dto *entities.RoleEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
