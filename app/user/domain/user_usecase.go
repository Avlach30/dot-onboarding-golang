package domain

import (
	"context"
	userdto "github.com/codespace-id/codespace-x/app/user/dto"
)

type Usecase interface {
	Create(ctx context.Context, dto userdto.RegisterRequest) error
	Profile(ctx context.Context, phoneNumber string) (res Entity, err error)
}
