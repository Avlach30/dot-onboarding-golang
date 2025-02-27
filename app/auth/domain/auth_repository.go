package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type AuthRepository interface {
	FindUserByEmailWithRoles(httpContext *gin.Context, email string) (*entities.UserEntity, error)
	FindUserByIDWithRoles(httpContext *gin.Context, id uuid.UUID) (*entities.UserEntity, error)
}
