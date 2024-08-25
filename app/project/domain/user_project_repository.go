package domain

import (
	"context"
	"database/sql"
)

type UserProjectRepository interface {
	CreateTx(ctx context.Context, dbTx *sql.Tx, userID, projectID int64) (err error)
}
