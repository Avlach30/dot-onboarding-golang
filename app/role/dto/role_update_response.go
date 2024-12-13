package dto

type RoleUpdateResponse struct {
	Name string `json:"name" binding:"required"`
	Key  string `json:"key" binding:"required"`
}
