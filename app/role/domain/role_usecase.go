package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/entities"
)

type RoleUsecase interface {
	Pagination(httpContext *gin.Context) ([]entities.RoleEntity, int)
	Create(httpContext *gin.Context, schema *entities.RoleEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID) *entities.RoleEntity
	Update(httpContext *gin.Context, id uuid.UUID, dto *entities.RoleEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
