package dto

type CreateProjectRequest struct {
	Name         string `json:"name"`
	Description  string `json:"description" validate:"required"`
	ServiceType  string `json:"service_type" validate:"required"`
	TimePriority string `json:"time_priority" validate:"required"`
}
