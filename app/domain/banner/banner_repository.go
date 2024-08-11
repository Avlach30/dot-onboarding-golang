package bannerdomain

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, skip, limit string) (res []Entity, err error)
}
