package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	roleEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type UserRepository interface {
	Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.UserEntity, int)
	Create(httpContext *gin.Context, payload *entities.UserEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.UserEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.UserEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsEmailExist(httpContext *gin.Context, email string) bool
	IsEmailExistExceptUserId(httpContext *gin.Context, email string, id uuid.UUID) bool
	FindRoleByIds(httpContext *gin.Context, ids []uuid.UUID) []roleEntities.RoleEntity
	DeleteUserRoles(httpContext *gin.Context, id uuid.UUID)
	IsExistById(httpContext *gin.Context, id uuid.UUID) bool
}
