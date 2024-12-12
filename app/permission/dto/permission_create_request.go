package dto

type PermissionCreateRequest struct {
	Name string `json:"name" validate:"required"`
	Key  string `json:"key" validate:"required"`
}
