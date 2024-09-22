package domain

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/user/userdomain"
)

type UserProjectRepository interface {
	CreateTx(ctx context.Context, dbTx *sql.Tx, userID, projectID int64) (err error)
	GetTalentInCharge(ctx context.Context, projectID int64, page, perPage int) (res []userdomain.Entity, err error)
}
