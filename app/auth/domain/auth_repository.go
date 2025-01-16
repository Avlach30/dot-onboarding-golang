package domain

import (
	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
)

type AuthRepository interface {
	FindUserByEmailWithRoles(httpContext *gin.Context, email string) (*domain.UserEntity, error)
}
