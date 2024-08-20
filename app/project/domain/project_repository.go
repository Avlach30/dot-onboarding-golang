package domain

import "context"

type Repository interface {
	Get(ctx context.Context, skip, limit string) (res []Entity, err error)
	Find(ctx context.Context, ID int) (res Entity, err error)
	Create(ctx context.Context, payload Entity) (err error)
	Update(ctx context.Context, payload Entity) (err error)
}
