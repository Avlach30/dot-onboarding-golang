package domain

import (
	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type AuthRepository interface {
	FindUserByEmailWithRoles(httpContext *gin.Context, email string) (*entities.UserEntity, error)
}
