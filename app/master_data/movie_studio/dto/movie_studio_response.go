package dto

import (
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
)

type MovieStudioIndexResponse struct {
	Id            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	ChairCapacity int       `json:"chair_capacity"`
	CreatedAt     time.Time    `json:"created_at"`
}

type MovieStudioDetailResponse struct {
	MovieStudioIndexResponse
	AdditionalCapacities []string `json:"additional_capacities"`
}

func NewMovieStudioIndexResponse(datas []entities.MovieStudioEntity) []MovieStudioIndexResponse {
	var movieStudios []MovieStudioIndexResponse

	for _, data := range datas {
		movieStudios = append(movieStudios, MovieStudioIndexResponse{
			Id:            data.ID,
			Name:          data.Name,
			ChairCapacity: data.ChairCapacity,
			CreatedAt:     data.CreatedAt,
		})
	}
	return movieStudios
}

func NewMovieStudioDetailResponse(movieStudio entities.MovieStudioEntity) MovieStudioDetailResponse {
	return MovieStudioDetailResponse{
		MovieStudioIndexResponse: MovieStudioIndexResponse{
			Id:            movieStudio.ID,
			Name:          movieStudio.Name,
			ChairCapacity: movieStudio.ChairCapacity,
			CreatedAt:     movieStudio.CreatedAt,
		},
		AdditionalCapacities: movieStudio.AdditionalCapacities,
	}
}
