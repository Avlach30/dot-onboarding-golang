package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
)

type MovieUsecase struct {
	movieRepository domain.MovieRepository
}

func NewMovieUsecase(movieRepository domain.MovieRepository) *MovieUsecase {
	return &MovieUsecase{
		movieRepository: movieRepository,
	}
}

func (u *MovieUsecase) Pagination(c *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieEntity, int) {
	movies, total := u.movieRepository.Pagination(queryDto, c)
	return movies, total
}

func (u *MovieUsecase) Create(c *gin.Context, payload *entities.MovieEntity) {
	u.movieRepository.Create(c, payload)
}

func (u *MovieUsecase) FindOneById(c *gin.Context, id uuid.UUID, trashed bool) *entities.MovieEntity {
	return u.movieRepository.FindOneById(c, id, trashed)
}

func (u *MovieUsecase) Update(c *gin.Context, id uuid.UUID, payload *entities.MovieEntity) {
	u.movieRepository.Update(c, id, payload)
}

func (u *MovieUsecase) Delete(c *gin.Context, id uuid.UUID) {
	u.movieRepository.Delete(c, id)
}