package repository

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/integration/domain"
	"gorm.io/gorm"
)

type LogIntegrationRepository struct {
	db *gorm.DB
}

// CreateLogIntegration implements domain.LogIntegrationRepository.
func (logIntegrationRepo *LogIntegrationRepository) CreateLogIntegration(payload *domain.LogIntegrationEntity) error {
	return logIntegrationRepo.db.Save(&payload).Error
}

func NewLogIntegrationRepository(db *gorm.DB) domain.LogIntegrationRepository {
	return &LogIntegrationRepository{
		db: db,
	}
}
