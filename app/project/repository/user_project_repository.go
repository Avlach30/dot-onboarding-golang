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
				project_id
			) 
		VALUES 
			(?, ?)
		`

	if _, err := dbTx.ExecContext(
		ctx,
		query,
		userID,
		projectID,
	); err != nil {
		return errors.Wrap(err, "UserProjectRepository.CreateTx.ExecContext")
	}

	return nil
}

func (r *UserProjectRepository) GetTalentInCharge(ctx context.Context, projectID int64, page, perPage int) (res []userdomain.Entity, err error) {
	query := `
		SELECT
		     u.fullname,
		     u.image_url,
		     up.project_role
		FROM
			user_projects up
		JOIN users u ON up.user_id = u.id
		WHERE up.project_id = ?
		AND up.is_project_owner = 0
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
		return res, errors.Wrap(err, "UserProjectRepository.GetTalentInCharge.QueryContext")
	}
	defer list.Close()

	for list.Next() {
		var user userdomain.Entity

		var projectRole sql.NullString
		var imageUrl = sql.NullString{}
		err = list.Scan(
			&user.Fullname,
			&imageUrl,
			&projectRole,
		)
		if err != nil {
			return res, errors.Wrap(err, "UserProjectRepository.GetTalentInCharge.QueryContext")
		}

		user.ImageURL = imageUrl.String
		user.Role = projectRole.String
		res = append(res, user)
	}

	return res, nil
}
