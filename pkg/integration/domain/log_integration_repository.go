package domain

type LogIntegrationRepository interface {
	CreateLogIntegration(payload *LogIntegrationEntity) error
}
