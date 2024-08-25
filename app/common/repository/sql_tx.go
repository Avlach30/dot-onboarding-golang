package commonrepo

import (
	"context"
	"database/sql"
)

type SqlTx interface {
	Begin(ctx context.Context) (*sql.Tx, error)
}

type sqlTransactor struct {
	connection *sql.DB
}

func NewSqlTx(db *sql.DB) SqlTx {
	return &sqlTransactor{
		connection: db,
	}
}

func (u *sqlTransactor) Begin(ctx context.Context) (*sql.Tx, error) {
	tx, err := u.connection.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
