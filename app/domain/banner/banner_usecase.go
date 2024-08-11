package bannerdomain

import "context"

type Usecase interface {
	Get(ctx context.Context, page, perPage int) (res []Entity, err error)
}
