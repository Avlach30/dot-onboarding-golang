package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type MovieRepository struct {
	model *gorm.DB
}

func NewMovieRepository(db *gorm.DB) domain.MovieRepository {
	return &MovieRepository{
		model: db.Model(&entities.MovieEntity{}),
	}
}

func (movie *MovieRepository) Pagination(queryDto *querydto.QueryDto, httpContext *gin.Context) ([]entities.MovieEntity, int) {
	query := movie.model.WithContext(httpContext)
	var movies []entities.MovieEntity
	var total int64

	// Query filter
	query = movie.queryFilter(query, queryDto)
	// Query sort
	query = movie.querySort(query, queryDto)

	
	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error counting movies: ", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(queryDto)).
		Find(&movies).Error
	if err != nil {
		log.Println("Error fetching movies: ", err)
		panic(*exception.ServerErrorException(err))
	}

	return movies, int(total)
}

func (movie *MovieRepository) queryFilter(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	if search := queryDto.Search; search != "" {
		query = query.Where("name ILIKE ?", search+"%")
	}

	return query
}

func (movie *MovieRepository) querySort(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	sortableColumns := []string{"name", "chair_capacity", "duration_in_minutes", "updated_at"}

	if sort := queryDto.SortBy; sort != "" {
		// Check if the sort column is valid
		if !utils.Contains(sortableColumns, sort) {
			panic(*exception.BussinessException("Invalid sort column"))
		}

		// Handle order query
		if order := queryDto.Order; order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order value"))
			}

			query = query.Order(sort + " " + order)
		}
	} else {
		query = query.Order("updated_at desc")
	}

	return query
}

func (movie *MovieRepository) Create(httpContext *gin.Context, payload *entities.MovieEntity) {
	movie.model = movie.model.WithContext(httpContext)

	err := movie.model.Create(payload).Error
	if err != nil {
		log.Println("Error creating movie:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movie *MovieRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieEntity {
	movie.model = movie.model.WithContext(httpContext)
	movieEntity := &entities.MovieEntity{}
	if trashed {
		movie.model = movie.model.Unscoped()
	}

	err := movie.model.
		First(&movieEntity, id).
		Error
	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Movie not found"))
	} else if err != nil {
		log.Println("Error movie find by id:", err)
		panic(*exception.ServerErrorException(err))
	}

	return movieEntity
}

func (movie *MovieRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieEntity) {
	movie.model = movie.model.WithContext(httpContext)

	err := movie.model.
		Where("id = ?", id).
		Updates(payload).
		Error
	if err != nil {
		log.Println("Error updating movie:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movie *MovieRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	movie.model = movie.model.WithContext(httpContext)

	err := movie.model.
		Delete(&entities.MovieEntity{}, id).
		Error
	if err != nil {
		log.Println("Error deleting movie:", err)
		panic(*exception.ServerErrorException(err))
	}
}

