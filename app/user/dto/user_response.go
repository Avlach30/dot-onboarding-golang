package dto

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type UserResponse struct {
	Id           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
}

func NewUserResponse(data entities.UserEntity) UserResponse {
	return UserResponse{
		Id:           data.ID,
		Email:        data.Email,
		Name:         data.Name,
	}
}