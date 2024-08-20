package bannerdomain

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, page, perPage int) (res []Entity, err error)
}
