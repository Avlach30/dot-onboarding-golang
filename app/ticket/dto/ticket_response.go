package dto

import (
	"time"

	"github.com/google/uuid"
	movieScheduleDto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_schedule/dto"
	userDto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type TicketIndexResponse struct {
	Id           uuid.UUID `json:"id"`
	Status       string    `json:"status"`
	SelectedChairs []string `json:"selected_chairs"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewTicketIndexResponse(datas []entities.TicketEntity) []TicketIndexResponse {
	formatedData := make([]TicketIndexResponse, len(datas))
	for i, data := range datas {
		formatedData[i] = TicketIndexResponse{
			Id:           data.ID,
			Status:       data.Status,
			SelectedChairs: data.SelectedChairs,
			UpdatedAt:    data.UpdatedAt,
		}
	}
	return formatedData
}

type TicketDetailResponse struct {
	TicketIndexResponse
	MovieSchedule movieScheduleDto.MovieScheduleDetailResponse `json:"movie_schedule"`
	User          userDto.UserResponse                  `json:"user"`
}

func NewTicketDetailResponse(data entities.TicketEntity) TicketDetailResponse {
	return TicketDetailResponse{
		TicketIndexResponse: TicketIndexResponse{
			Id:           data.ID,
			Status:       data.Status,
			SelectedChairs: data.SelectedChairs,
			UpdatedAt:    data.UpdatedAt,
		},
		MovieSchedule: movieScheduleDto.NewMovieScheduleDetailResponse(data.MovieSchedule),
		User:          userDto.NewUserResponse(data.User),
	}
}

