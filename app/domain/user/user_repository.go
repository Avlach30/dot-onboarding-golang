package userdomain

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, payload Entity) error
}
