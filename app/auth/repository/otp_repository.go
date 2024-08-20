package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/auth/domain"
	"github.com/pkg/errors"
)

type OtpRepository struct {
	db *sql.DB
}

func NewOtpRepository(db *sql.DB) domain.OtpRepository {
	return &OtpRepository{
		db: db,
	}
}

// Create implements authdomain.OtpRepository.
func (r *OtpRepository) Create(ctx context.Context, payload domain.OtpEntity) error {
	query := `
		INSERT INTO 
			otps(
				code, 
				identifier, 
				trial, 
			   is_valid,
			   expired_at
			) 
		VALUES 
			(?, ?, ?, ?, ?)
		`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		payload.Code,
		payload.Identifier,
		payload.Trial,
		payload.IsValid,
		payload.ExpiredAt,
	); err != nil {
		return errors.Wrap(err, "OtpRepository.Create.ExecContext")
	}

	return nil
}

// FindByIdentifier implements authdomain.OtpRepository.
func (r *OtpRepository) FindByIdentifier(ctx context.Context, identifier string) (res domain.OtpEntity, err error) {
	query := `
		SELECT
			code, 
			identifier, 
			trial, 
			is_valid,
			expired_at
		FROM
			otps
		WHERE
			identifier = ?
		`

	if err := r.db.QueryRowContext(
		ctx,
		query,
		identifier,
	).Scan(
		&res.Code,
		&res.Identifier,
		&res.Trial,
		&res.IsValid,
		&res.ExpiredAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "OtpRepository.FindByIdentifier.QueryRowContext")
	}

	return res, nil
}

// Update implements authdomain.OtpRepository.
func (r *OtpRepository) UpdateByIdentifier(ctx context.Context, identifier string, payload domain.OtpEntity) error {
	query := `
		UPDATE 
			otps
		SET 
			code = ?, 
			identifier = ?, 
			trial = ?, 
			is_valid = ?, 
			expired_at = ?
		WHERE 
			identifier = ?;
		`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		payload.Code,
		payload.Identifier,
		payload.Trial,
		payload.IsValid,
		payload.ExpiredAt,

		identifier,
	); err != nil {
		return errors.Wrap(err, "OtpRepository.UpdateByIdentifier.ExecContext")
	}

	return nil
}

// Upsert implements authdomain.OtpRepository.
func (r *OtpRepository) Upsert(ctx context.Context, payload domain.OtpEntity) error {
	query := `
		INSERT INTO otps (identifier, code, trial, is_valid, expired_at)
        VALUES (?, ?, ?, ?, ?)
        ON DUPLICATE KEY UPDATE
            code = VALUES(code),
            trial = VALUES(trial),
            is_valid = VALUES(is_valid),
            expired_at = VALUES(expired_at)
		`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		payload.Identifier,
		payload.Code,
		payload.Trial,
		payload.IsValid,
		payload.ExpiredAt,
	); err != nil {
		return errors.Wrap(err, "OtpRepository.Upsert.ExecContext")
	}

	return nil
}
