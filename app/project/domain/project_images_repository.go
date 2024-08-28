package domain

import (
	"context"
	"database/sql"
)

type ProjectImagesRepository interface {
	CreateTx(ctx context.Context, dbTx *sql.Tx, imageUrl string, projectID int64) (err error)
}
