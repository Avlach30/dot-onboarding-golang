package domain

import (
	"time"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
)

type AuthUsecase interface {
	SignInBasic(email string, password string) (token string, expirationTime time.Time)
	SignInLDAP(username string, password string) (token string, expirationTime time.Time)
	SignInByOIDCCode(code string) (token string, expirationTime time.Time)
	CreateJWTToken(user *domain.UserEntity) (token string, expirationTime time.Time)
}
