package domain

import (
	"context"
	"database/sql"
)

type Repository interface {
	Get(ctx context.Context, page, perPage int) (res []Entity, err error)
	Find(ctx context.Context, UUID string) (res Entity, err error)
	CreateTx(ctx context.Context, dbTx *sql.Tx, payload Entity) (res Entity, err error)
	Update(ctx context.Context, payload Entity) (err error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string, page, perPage int) (res []Entity, err error)
	GetByStatus(ctx context.Context, page, perPage int, status string) (res []Entity, err error)
}
