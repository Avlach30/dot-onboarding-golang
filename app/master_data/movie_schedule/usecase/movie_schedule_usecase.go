package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_schedule/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type MovieScheduleUsecase struct {
	movieScheduleRepository domain.MovieScheduleRepository
}

func NewMovieScheduleUsecase(movieScheduleRepository domain.MovieScheduleRepository) *MovieScheduleUsecase {
	return &MovieScheduleUsecase{
		movieScheduleRepository: movieScheduleRepository,
	}
}

func (movieScheduleUsecase *MovieScheduleUsecase) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieScheduleEntity, int) {
	return movieScheduleUsecase.movieScheduleRepository.Pagination(httpContext, queryDto)
}

func (movieScheduleUsecase *MovieScheduleUsecase) Create(httpContext *gin.Context, payload *entities.MovieScheduleEntity) {
	movieScheduleUsecase.movieScheduleRepository.Create(httpContext, payload)
}

func (movieScheduleUsecase *MovieScheduleUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieScheduleEntity {
	return movieScheduleUsecase.movieScheduleRepository.FindOneById(httpContext, id, trashed)
}

func (movieScheduleUsecase *MovieScheduleUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieScheduleEntity) {
	movieScheduleUsecase.movieScheduleRepository.Update(httpContext, id, payload)
}

func (movieScheduleUsecase *MovieScheduleUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	movieScheduleUsecase.movieScheduleRepository.Delete(httpContext, id)
}