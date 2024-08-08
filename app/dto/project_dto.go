package dto

type ListProjectResponse struct {
	UUID        string         `json:"uuid"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	ServiceType string         `json:"service_type"`
	Status      string         `json:"status"`
	CreatedAt   string         `json:"created_at"`
	Astrodevs   []UserResponse `json:"astrodevs"`
}
