package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
)

type UserRepository interface {
	Pagination(httpContext *gin.Context) ([]UserEntity, int)
	Create(httpContext *gin.Context, payload *UserEntity)
	FindById(httpContext *gin.Context, id uuid.UUID, trashed bool) *UserEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *UserEntity)
	Delete(httpContext *gin.Context, id uuid.UUID)
	IsEmailExist(httpContext *gin.Context, email string) bool
	IsEmailExistExceptUserId(httpContext *gin.Context, email string, id uuid.UUID) bool
	FindRoleByIds(httpContext *gin.Context, ids []uuid.UUID) []domain.RoleEntity
	DeleteUserRoles(httpContext *gin.Context, id uuid.UUID)
}
