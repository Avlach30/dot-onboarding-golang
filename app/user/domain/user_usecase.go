package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserUsecase interface {
	Pagination(httpContext *gin.Context) ([]UserEntity, int)
	Create(httpContext *gin.Context, payload *UserEntity, roleIds []uuid.UUID)
	FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *UserEntity
	Update(httpContext *gin.Context, id uuid.UUID, payload *UserEntity, roleIds []uuid.UUID)
	Delete(httpContext *gin.Context, id uuid.UUID)
}
