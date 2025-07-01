package dto

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type TicketUpdateRequest struct {
	UserId          uuid.UUID `json:"user_id" binding:"required,uuid"`
	MovieScheduleId uuid.UUID `json:"movie_schedule_id" binding:"required,uuid"`
	Status          string    `json:"status" binding:"required"`
	SelectedChairs  []string  `json:"selected_chairs" binding:"required"`
}

func AssignUpdate(request *TicketUpdateRequest) entities.TicketEntity {
	return entities.TicketEntity{
		UserId:          request.UserId,
		MovieScheduleId: request.MovieScheduleId,
		Status:          request.Status,
		SelectedChairs:  request.SelectedChairs,
	}
}
