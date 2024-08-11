package repository

import (
	"context"
	"database/sql"

	bannerdomain "github.com/codespace-id/codespace-x/app/domain/banner"
)

type BannernRepository struct {
	db *sql.DB
}

func NewBannerRepository(db *sql.DB) bannerdomain.Repository {
	return &BannernRepository{
		db: db,
	}
}

// Get implements bannerdomain.Repository.
func (b *BannernRepository) Get(ctx context.Context, skip string, limit string) (res []bannerdomain.Entity, err error) {
	panic("unimplemented")
}
