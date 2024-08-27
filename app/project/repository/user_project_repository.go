package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/codespace-id/codespace-x/app/user/userdomain"
	"github.com/codespace-id/codespace-x/pkg/utils"
	"github.com/pkg/errors"
)

type UserProjectRepository struct {
	db *sql.DB
}

func NewUserProjectRepository(db *sql.DB) domain.UserProjectRepository {
	return &UserProjectRepository{
		db: db,
	}
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

func (r *UserProjectRepository) GetByProjectID(ctx context.Context, projectID int64, page, perPage int) (res []userdomain.Entity, err error) {
	query := `
		SELECT
		     u.fullname,
		     u.image_url,
		     up.project_role
		FROM
			user_projects up
		JOIN users u ON up.user_id = u.id
		WHERE up.project_id = ?
		LIMIT ? OFFSET ?
		`

	list, err := r.db.QueryContext(
		ctx,
		query,
		projectID,
		perPage,
		utils.GetPaginationOffset(page, perPage),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "UserProjectRepository.GetByProjectID.QueryContext")
	}
	defer list.Close()

	for list.Next() {
		var user userdomain.Entity

		err = list.Scan(
			&user.Fullname,
			&user.ImageURL,
			&user.Role,
		)
		if err != nil {
			return res, errors.Wrap(err, "UserProjectRepository.GetByProjectID.QueryContext")
		}

		res = append(res, user)
	}

	return res, nil
}
