package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/user/userdomain"

	"github.com/pkg/errors"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) userdomain.Repository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, payload userdomain.Entity) error {

	query := `
		INSERT INTO 
			users(
				fullname, 
				identity_number, 
				phone_number, 
			    gender,
			    password
			) 
		VALUES 
			(?, ?, ?, ?, ?)
		`

	if _, err := r.db.ExecContext(
		ctx,
		query,
		payload.Fullname,
		payload.IdentityNumber,
		payload.PhoneNumber,
		payload.Gender,
		payload.Password,
	); err != nil {
		return errors.Wrap(err, "UserRepository.CreateTx.ExecContext")
	}

	return nil
}

func (r *UserRepository) Find(ctx context.Context, phoneNumber string) (res userdomain.Entity, err error) {

	query := `
		SELECT
		    id,
			fullname, 
			identity_number, 
			phone_number, 
		   gender
		FROM
			users
		WHERE
			phone_number = ?
		`

	if err := r.db.QueryRowContext(
		ctx,
		query,
		phoneNumber,
	).Scan(
		&res.ID,
		&res.Fullname,
		&res.IdentityNumber,
		&res.PhoneNumber,
		&res.Gender,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return res, nil
		}
		return res, errors.Wrap(err, "UserRepository.CreateTx.ExecContext")
	}

	return res, nil
}
