package dto

import (
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type TicketCreateRequest struct {
	UserId          uuid.UUID `json:"user_id" binding:"required,uuid"`
	MovieScheduleId uuid.UUID `json:"movie_schedule_id" binding:"required,uuid"`
	SelectedChairs  []string  `json:"selected_chairs" binding:"required"`
}

func AssignCreate(request *TicketCreateRequest) entities.TicketEntity {
	return entities.TicketEntity{
		UserId:          request.UserId,
		MovieScheduleId: request.MovieScheduleId,
		SelectedChairs:  request.SelectedChairs,
		Status:          "confirmed",
	}
}
