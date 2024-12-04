package userdomain

import (
	"context"

	userdto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
)

type Usecase interface {
	Create(ctx context.Context, dto userdto.RegisterRequest) error
	Profile(ctx context.Context, phoneNumber string) (res Entity, err error)
	Delete(ctx context.Context, phoneNumber string) error
}
