package dto

import (
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type MovieIndexResponse struct {
	Id                uuid.UUID `json:"id"`
	Title             string    `json:"title"`
	Genre             string    `json:"genre"`
	DurationInMinutes int       `json:"duration_in_minutes"`
	UpdatedAt         time.Time    `json:"updated_at"`
}

type MovieDetailResponse struct {
	MovieIndexResponse
	PosterUrl         string    `json:"poster_url"`
	Description       string    `json:"description"`
}

func NewMovieIndexResponse(datas []entities.MovieEntity) []MovieIndexResponse {
	var movies []MovieIndexResponse

	for _, data := range datas {
		movies = append(movies, MovieIndexResponse{
			Id:                data.ID,
			Title:             data.Title,
			Genre:             data.Genre,
			DurationInMinutes: data.DurationInMinutes,
			UpdatedAt:         data.UpdatedAt,
		})
	}
	return movies
}

func NewMovieDetailResponse(movie entities.MovieEntity) MovieDetailResponse {
	return MovieDetailResponse{
		MovieIndexResponse: MovieIndexResponse{
			Id:                movie.ID,
			Title:             movie.Title,
			Genre:             movie.Genre,
			DurationInMinutes: movie.DurationInMinutes,
			UpdatedAt:         movie.UpdatedAt,
		},
		PosterUrl:         movie.PosterUrl,
		Description:       movie.Desciption,
	}
}