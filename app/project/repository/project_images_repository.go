package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/pkg/errors"
)

type ProjectImagesRepository struct {
	db *sql.DB
}

func NewProjectImagesRepository(db *sql.DB) domain.ProjectImagesRepository {
	return &ProjectImagesRepository{
		db: db,
	}
}

func (r *ProjectImagesRepository) CreateTx(ctx context.Context, dbTx *sql.Tx, imageUrl string, projectID int64) (err error) {
	query := `
		INSERT INTO 
			user_projects(
				image_url, 
				project_id, 
				is_thumbnail
			) 
		VALUES 
			(?, ?, ?)
		`

	if _, err := dbTx.ExecContext(
		ctx,
		query,
		imageUrl,
		projectID,
		1,
	); err != nil {
		return errors.Wrap(err, "ProjectImagesRepository.CreateTx.ExecContext")
	}

	return nil
}
