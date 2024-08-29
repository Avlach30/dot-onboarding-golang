package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/codespace-id/codespace-x/pkg/utils"
	"github.com/pkg/errors"
)

type ProjectHistoryRepository struct {
	db *sql.DB
}

func NewProjectHistoryRepository(db *sql.DB) domain.ProjectHistoryRepository {
	return &ProjectHistoryRepository{
		db: db,
	}
}
func (r *ProjectHistoryRepository) Get(ctx context.Context, projectUUID string, page, perPage int) (res []domain.ProjectHistoryEntity, err error) {
	query := `
		SELECT
		    ph.id,
			ph.title,
			ph.description, 
			ph.history_type,
			ph.attachment_url,
			ph.attachment_title,
			ph.created_at
		FROM
			project_histories ph
		JOIN projects ON projects.id = ph.project_id AND projects.uuid = ?
		ORDER BY ph.created_at DESC
		LIMIT ? OFFSET ?
		`

	list, err := r.db.QueryContext(
		ctx,
		query,
		projectUUID,
		perPage,
		utils.GetPaginationOffset(page, perPage),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "ProjectHistoryRepository.Get.QueryContext")
	}
	defer list.Close()

	for list.Next() {
		var project domain.ProjectHistoryEntity

		var attachmentUrl sql.NullString
		var attachmentTitle sql.NullString
		err = list.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&project.HistoryType,
			&attachmentUrl,
			&attachmentTitle,
			&project.CreatedAt,
		)
		if err != nil {
			return res, errors.Wrap(err, "ProjectHistoryRepository.Get.QueryContext")
		}

		project.AttachmentUrl = attachmentUrl.String
		project.AttachmentTitle = attachmentTitle.String
		res = append(res, project)
	}

	return res, nil
}
