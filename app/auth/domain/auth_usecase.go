package domain

import (
	"context"
	"time"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
)

type AuthUsecase interface {
	SignInBasic(context *context.Context, email string, password string) (token string, expirationTime time.Time)
	SignInLDAP(context *context.Context, username string, password string) (token string, expirationTime time.Time)
	SignInByOIDCCode(context *context.Context, code string) (token string, expirationTime time.Time)
	CreateJWTToken(user *domain.UserEntity) (token string, expirationTime time.Time)
}
