package domain

import (
	"context"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
)

type AuthRepository interface {
	FindUserByEmail(context *context.Context, email string) (*domain.UserEntity, error)
}
