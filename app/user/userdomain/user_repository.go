package userdomain

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, payload Entity) error
	Find(ctx context.Context, phoneNumber string) (res Entity, err error)
}
