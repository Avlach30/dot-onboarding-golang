package dto

import "github.com/google/uuid"

type UserCreateRequest struct {
	Name     string      `json:"name" binding:"required"`
	Email    string      `json:"email" binding:"required,email"`
	Password string      `json:"password" binding:"required"`
	RoleIds  []uuid.UUID `json:"role_ids" binding:"required,dive,uuid"`
}
