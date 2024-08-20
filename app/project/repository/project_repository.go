package repository

import (
	"context"
	"database/sql"
	"github.com/codespace-id/codespace-x/app/project/domain"
)

type ProjectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) domain.Repository {
	return &ProjectRepository{
		db: db,
	}
}

// Create implements projectdomain.Repository.
func (p *ProjectRepository) Create(ctx context.Context, payload domain.Entity) (err error) {
	panic("unimplemented")
}

// Find implements projectdomain.Repository.
func (p *ProjectRepository) Find(ctx context.Context, ID int) (res domain.Entity, err error) {
	panic("unimplemented")
}

// Get implements projectdomain.Repository.
func (p *ProjectRepository) Get(ctx context.Context, skip string, limit string) (res []domain.Entity, err error) {
	panic("unimplemented")
}

// Update implements projectdomain.Repository.
func (p *ProjectRepository) Update(ctx context.Context, payload domain.Entity) (err error) {
	panic("unimplemented")
}
