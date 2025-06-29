package dto

import (
	"time"

	"github.com/google/uuid"
	movieDto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie/dto"
	movieStudioDto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_studio/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type MovieScheduleIndexResponse struct {
	Id           uuid.UUID `json:"id"`
	ShowDatetime time.Time `json:"show_datetime"`
	Price        float64   `json:"price"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MovieScheduleDetailResponse struct {
	MovieScheduleIndexResponse
	MovieID       uuid.UUID                               `json:"movie_id"`
	Movie         movieDto.MovieDetailResponse             `json:"movie"`
	MovieStudioID uuid.UUID                               `json:"movie_studio_id"`
	MovieStudio   movieStudioDto.MovieStudioDetailResponse `json:"movie_studio"`
}

func NewMovieScheduleIndexResponse(datas []MovieScheduleIndexResponse) []MovieScheduleIndexResponse {
	var movieSchedules []MovieScheduleIndexResponse

	for _, data := range datas {
		movieSchedules = append(movieSchedules, MovieScheduleIndexResponse{
			Id:           data.Id,
			ShowDatetime: data.ShowDatetime,
			Price:        data.Price,
			UpdatedAt:    data.UpdatedAt,
		})
	}
	return movieSchedules
}

func NewMovieScheduleDetailResponse(movieSchedule entities.MovieScheduleEntity) MovieScheduleDetailResponse {
	return MovieScheduleDetailResponse{
		MovieScheduleIndexResponse: MovieScheduleIndexResponse{
			Id:           movieSchedule.ID,
			ShowDatetime: movieSchedule.ShowDatetime,
			Price:        movieSchedule.Price,
			UpdatedAt:    movieSchedule.UpdatedAt,
		},
		Movie:       movieDto.NewMovieDetailResponse(movieSchedule.Movie),
		MovieStudio: movieStudioDto.NewMovieStudioDetailResponse(movieSchedule.MovieStudio),
	}
}
