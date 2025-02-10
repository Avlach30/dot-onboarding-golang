package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	roleEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/entities"
)

type UserRepository interface {
	Pagination(httpContext *gin.Context) ([]entities.UserEntity, int)
	Create(httpContext *gin.Context, payload *entities.UserEntity)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.UserEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *entities.UserEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsEmailExist(httpContext *gin.Context, email string) bool
	IsEmailExistExceptUserId(httpContext *gin.Context, email string, id uuid.UUID) bool
	FindRoleByIds(httpContext *gin.Context, ids []uuid.UUID) []roleEntities.RoleEntity
	DeleteUserRoles(httpContext *gin.Context, id uuid.UUID)
}
