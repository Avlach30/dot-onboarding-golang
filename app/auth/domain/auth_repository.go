package domain

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
)

type AuthRepository interface {
	FindUserByEmail(email string) (*domain.UserEntity, error)
}
