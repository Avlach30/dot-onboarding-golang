package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type RoleRepository interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.RoleEntity, int)
	Create(httpContext *gin.Context, payload *entities.RoleEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.RoleEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.RoleEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsKeyExist(httpContext *gin.Context, key string) bool
	IsKeyExistExceptRoleId(httpContext *gin.Context, key string, id uuid.UUID) bool
}
