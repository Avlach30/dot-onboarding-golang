package dto

type PermissionUpdateRequest struct {
	Name string `json:"name" validate:"required"`
	Key  string `json:"key" validate:"required"`
}
