package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type MovieStudioRepository interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieStudioEntity, int)
	Create(httpContext *gin.Context, payload *entities.MovieStudioEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieStudioEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieStudioEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsExistsByName(httpContext *gin.Context, name string, id *uuid.UUID) bool
}