package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/pkg/errors"
)

type UserProjectRepository struct {
}

func NewUserProjectRepository() domain.UserProjectRepository {
	return &UserProjectRepository{}
}

func (r *UserProjectRepository) CreateTx(ctx context.Context, dbTx *sql.Tx, userID, projectID int64) (err error) {
	query := `
		INSERT INTO 
			user_projects(
				user_id, 
				project_id, 
				is_project_owner
			) 
		VALUES 
			(?, ?, ?)
		`

	if _, err := dbTx.ExecContext(
		ctx,
		query,
		userID,
		projectID,
		1,
	); err != nil {
		return errors.Wrap(err, "UserProjectRepository.CreateTx.ExecContext")
	}

	return nil
}
