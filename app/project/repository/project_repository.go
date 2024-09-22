package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/project/domain"
	"github.com/codespace-id/codespace-x/pkg/common/enum"
	"github.com/codespace-id/codespace-x/pkg/utils"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) domain.Repository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) CreateTx(ctx context.Context, dbTx *sql.Tx, payload domain.Entity) (res domain.Entity, err error) {

	newUUID := uuid.New().String()
	query := `
		INSERT INTO 
			projects(
				uuid, 
				name, 
				description,
				service_type,
				status,
				budget,
				target_time
			) 
		VALUES 
			(?, ?, ?, ?, ?, ?, ?)
		`

	data, err := dbTx.ExecContext(
		ctx,
		query,
		newUUID,
		payload.Name,
		payload.Description,
		payload.ServiceType,
		enum.INQUIRY.Value(),
		0,
		payload.TargetTime,
	)
	if err != nil {
		return res, errors.Wrap(err, "ProjectRepository.CreateTx.ExecContext")
	}

	projectID, _ := data.LastInsertId()

	res.UUID = newUUID
	res.ID = projectID

	return res, nil
}

func (r *ProjectRepository) Find(ctx context.Context, UUID string) (res domain.Entity, err error) {
	query := `
		SELECT
		    p.id AS id,
		   p.uuid,
			p.name,
			p.description, 
			p.service_type,
			p.status,
			p.created_at,
			pi.image_url AS thumbnail_image_url,
			p.target_time
		FROM
			projects p
		JOIN project_images pi ON p.id = pi.project_id AND pi.is_thumbnail = 1
		WHERE p.uuid = ?
		`

	var project domain.Entity
	var thumbnailImageURL sql.NullString
	var targetTime sql.NullString

	if err := r.db.QueryRowContext(
		ctx,
		query,
		UUID,
	).Scan(
		&project.ID,
		&project.UUID,
		&project.Name,
		&project.Description,
		&project.ServiceType,
		&project.Status,
		&project.CreatedAt,
		&thumbnailImageURL,
		&targetTime,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "ProjectRepository.Find.QueryRowContext")
	}

	project.ThumbnailImageURL = thumbnailImageURL.String
	project.TargetTime = targetTime.String

	return project, nil
}

func (r *ProjectRepository) Get(ctx context.Context, page, perPage int) (res []domain.Entity, err error) {
	query := `
		SELECT
		    uuid,
			name,
			description, 
			service_type,
			status,
			p.created_at,
			MAX(pi.image_url) AS thumbnail_image_url,
			astrodevs
		FROM
			projects p
		LEFT JOIN project_images pi ON p.id = pi.project_id AND pi.is_thumbnail = 1
		WHERE p.deleted_at IS NULL
		GROUP BY uuid, name, description, service_type, status,astrodevs, created_at
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
		return res, errors.Wrap(err, "ProjectRepository.Get.QueryContext")
	}
	defer list.Close()

	for list.Next() {
		var project domain.Entity

		var thumbnailImageURL sql.NullString
		var astrodevs sql.NullString

		err = list.Scan(
			&project.UUID,
			&project.Name,
			&project.Description,
			&project.ServiceType,
			&project.Status,
			&project.CreatedAt,
			&thumbnailImageURL,
			&astrodevs,
		)
		if err != nil {
			return res, errors.Wrap(err, "ProjectRepository.Get.QueryContext")
		}

		project.ThumbnailImageURL = thumbnailImageURL.String
		project.Astrodevs = astrodevs.String

		res = append(res, project)
	}

	return res, nil
}

func (r *ProjectRepository) Update(ctx context.Context, payload domain.Entity) (err error) {
	panic("unimplemented")
}

func (r *ProjectRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string, page, perPage int) (res []domain.Entity, err error) {
	query := `
		SELECT
		    p.id as id,
		     uuid,
			name,
			description, 
			service_type,
			status,
			p.created_at,
			MAX(pi.image_url) AS thumbnail_image_url
		FROM
			projects p
		JOIN user_projects up ON p.id = up.project_id
		JOIN users u ON up.user_id = u.id
		LEFT JOIN project_images pi ON p.id = pi.project_id AND pi.is_thumbnail = 1
		WHERE u.phone_number = ? AND p.deleted_at = NULL
		GROUP BY id, uuid, name, description, service_type, status, created_at
		LIMIT ? OFFSET ?
		`

	list, err := r.db.QueryContext(
		ctx,
		query,
		phoneNumber,
		perPage,
		utils.GetPaginationOffset(page, perPage),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "ProjectRepository.Get.QueryContext")
	}
	defer list.Close()

	for list.Next() {
		var project domain.Entity

		var thumbnailImageURL sql.NullString
		err = list.Scan(
			&project.ID,
			&project.UUID,
			&project.Name,
			&project.Description,
			&project.ServiceType,
			&project.Status,
			&project.CreatedAt,
			&thumbnailImageURL,
		)
		if err != nil {
			return res, errors.Wrap(err, "ProjectRepository.Get.QueryContext")
		}

		project.ThumbnailImageURL = thumbnailImageURL.String
		res = append(res, project)
	}

	return res, nil
}

func (r *ProjectRepository) GetByStatus(ctx context.Context, page, perPage int, status string) (res []domain.Entity, err error) {
	query := `
		SELECT
		    uuid,
			name,
			description, 
			service_type,
			status,
			p.created_at,
			MAX(pi.image_url) AS thumbnail_image_url
		FROM
			projects p
		LEFT JOIN project_images pi ON p.id = pi.project_id AND pi.is_thumbnail = 1
		WHERE p.status = ?
		GROUP BY uuid, name, description, service_type, status, created_at
		LIMIT ? OFFSET ?
		`

	list, err := r.db.QueryContext(
		ctx,
		query,
		status,
		perPage,
		utils.GetPaginationOffset(page, perPage),
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "ProjectRepository.GetByStatus.QueryContext")
	}
	defer list.Close()

	for list.Next() {
		var project domain.Entity

		var thumbnailImageURL sql.NullString
		err = list.Scan(
			&project.UUID,
			&project.Name,
			&project.Description,
			&project.ServiceType,
			&project.Status,
			&project.CreatedAt,
			&thumbnailImageURL,
		)
		if err != nil {
			return res, errors.Wrap(err, "ProjectRepository.GetByStatus.QueryContext")
		}

		project.ThumbnailImageURL = thumbnailImageURL.String
		res = append(res, project)
	}

	return res, nil
}
