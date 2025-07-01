package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_studio/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type MovieStudioUsecase struct {
	movieStudioRepo domain.MovieStudioRepository
}

func NewMovieStudioUsecase(movieStudioRepo domain.MovieStudioRepository) domain.MovieStudioUsecase {
	return &MovieStudioUsecase{
		movieStudioRepo: movieStudioRepo,
	}
}

func (movieStudioUsecase *MovieStudioUsecase) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieStudioEntity, int) {
	movieStudios, total := movieStudioUsecase.movieStudioRepo.Pagination(httpContext, queryDto)

	return movieStudios, total
}

func (movieStudioUsecase *MovieStudioUsecase) Create(httpContext *gin.Context, payload *entities.MovieStudioEntity) {
	movieStudioUsecase.movieStudioRepo.Create(httpContext, payload)
}

func (movieStudioUsecase *MovieStudioUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieStudioEntity {
	movieStudio := movieStudioUsecase.movieStudioRepo.FindOneById(httpContext, id, trashed)

	return movieStudio
}

func (movieStudioUsecase *MovieStudioUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieStudioEntity) {
	movieStudioUsecase.movieStudioRepo.Update(httpContext, id, payload)
}

func (movieStudioUsecase *MovieStudioUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	movieStudioUsecase.movieStudioRepo.Delete(httpContext, id)
}