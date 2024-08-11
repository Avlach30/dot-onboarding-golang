package repository

import (
	"context"
	"database/sql"

	projectdomain "github.com/codespace-id/codespace-x/app/domain/project"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) projectdomain.Repository {
	return &ProjectRepository{
		db: db,
	}
}

// Create implements projectdomain.Repository.
func (p *ProjectRepository) Create(ctx context.Context, payload projectdomain.Entity) (err error) {
	panic("unimplemented")
}

// Find implements projectdomain.Repository.
func (p *ProjectRepository) Find(ctx context.Context, ID int) (res projectdomain.Entity, err error) {
	panic("unimplemented")
}

// Get implements projectdomain.Repository.
func (p *ProjectRepository) Get(ctx context.Context, skip string, limit string) (res []projectdomain.Entity, err error) {
	panic("unimplemented")
}

// Update implements projectdomain.Repository.
func (p *ProjectRepository) Update(ctx context.Context, payload projectdomain.Entity) (err error) {
	panic("unimplemented")
}
