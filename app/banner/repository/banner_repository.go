package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/banner/bannerdomain"
	"github.com/codespace-id/codespace-x/pkg/utils"
	"github.com/pkg/errors"
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
func (r *BannernRepository) Get(ctx context.Context, page int, perPage int) (res []bannerdomain.Entity, err error) {
	query := `
		SELECT
			title, 
			description, 
			image_url
		FROM
			banners
		LIMIT ? OFFSET ?
		`

	list, err := r.db.QueryContext(
		ctx,
		query,
		perPage,
		utils.GetPaginationOffset(page, perPage),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "UserRepository.Create.ExecContext")
	}
	defer list.Close()

	for list.Next() {
		var banner bannerdomain.Entity

		err = list.Scan(&banner.Title, &banner.Description, &banner.ImageURL)
		if err != nil {
			return res, errors.Wrap(err, "UserRepository.Create.Scan")
		}

		res = append(res, banner)
	}

	return res, nil
}
