package domain

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type AuthUsecase interface {
	SignInBasic(httpContext *gin.Context, email string, password string) (token string, expirationTime time.Time)
	SignInLDAP(httpContext *gin.Context, username string, password string) (token string, expirationTime time.Time)
	SignInByOIDCCode(httpContext *gin.Context, code string, redirectUri string) (token string, expirationTime time.Time)
	CreateJWTToken(user *entities.UserEntity) (token string, expirationTime time.Time)
}
