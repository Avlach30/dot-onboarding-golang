package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/domain/user"
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
		return errors.Wrap(err, "UserRepository.Create.ExecContext")
	}

	return nil
}
