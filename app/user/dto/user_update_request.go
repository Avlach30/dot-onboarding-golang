package dto

import "github.com/google/uuid"

type UserUpdateRequest struct {
	Name    string      `json:"name" binding:"required"`
	Email   string      `json:"email" binding:"required,email"`
	RoleIds []uuid.UUID `json:"role_ids" binding:"required,dive,uuid,min=1"`
}
