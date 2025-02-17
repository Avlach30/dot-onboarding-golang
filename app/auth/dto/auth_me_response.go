package dto

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type AuthMeResponse struct {
	ID    uuid.UUID          `json:"id" binding:"required"`
	Name  string             `json:"name" binding:"required"`
	Email string             `json:"email" binding:"required"`
	Roles []dto.RoleResponse `json:"roles" binding:"required,dive"`

	// Roles []RoleResponse `json:"roles" binding:"required,dive"`
}

func AuthMeResponseFromEntity(data entities.UserEntity) AuthMeResponse {
	return AuthMeResponse{
		ID:    data.ID,
		Name:  data.Name,
		Email: data.Email,
		Roles: dto.RoleResponseFromEntities(data.Roles),
	}
}
