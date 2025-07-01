package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_schedule/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type MovieScheduleRepository struct {
	model *gorm.DB
}

func NewMovieScheduleRepository(db *gorm.DB) domain.MovieScheduleRepository {
	return &MovieScheduleRepository{
		model: db.Model(&entities.MovieScheduleEntity{}),
	}
}

func (movieSchedule *MovieScheduleRepository) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieScheduleEntity, int) {
	query := movieSchedule.model.WithContext(httpContext)
	var movieSchedules []entities.MovieScheduleEntity
	var total int64

	// Query filter
	query = movieSchedule.queryFilter(query, queryDto)
	// Query sort
	query = movieSchedule.querySort(query, queryDto)

	err := query.Count(&total).
		Error
	if err != nil {
		log.Println("Error counting movie schedules:", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(queryDto)).
		Find(&movieSchedules).Error

	if err != nil {
		log.Println("Error fetching movie schedules:", err)
		panic(*exception.ServerErrorException(err))
	}

	return movieSchedules, int(total)
}

// Func filter for pagination
func (movieSchedule *MovieScheduleRepository) queryFilter(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	if search := queryDto.Search; search != "" {
		query = query.Where("show_datetime LIKE ?", search+"%")
	}

	return query
}

func (movieSchedule *MovieScheduleRepository) querySort(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	sortableColumns := []string{"show_datetime", "price", "updated_at"}

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

func (movieSchedule *MovieScheduleRepository) Create(httpContext *gin.Context, payload *entities.MovieScheduleEntity) {
	movieSchedule.model = movieSchedule.model.WithContext(httpContext)

	err := movieSchedule.model.Create(payload).Error
	if err != nil {
		log.Println("Error creating movie schedule:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movieSchedule *MovieScheduleRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieScheduleEntity {
	movieSchedule.model = movieSchedule.model.WithContext(httpContext)
	movieScheduleEntity := &entities.MovieScheduleEntity{}
	if trashed {
		movieSchedule.model = movieSchedule.model.Unscoped()
	}

	err := movieSchedule.model.
		Preload("Movie").
		Preload("MovieStudio").
		First(&movieScheduleEntity, id).
		Error
	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Movie Schedule not found"))
	} else if err != nil {
		log.Println("Error movie schedule find by id:", err)
		panic(*exception.ServerErrorException(err))
	}

	return movieScheduleEntity
}

func (movieSchedule *MovieScheduleRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieScheduleEntity) {
	movieSchedule.model = movieSchedule.model.WithContext(httpContext)

	err := movieSchedule.model.
		Where("id = ?", id).
		Updates(payload).
		Error
	if err != nil {
		log.Println("Error updating movie schedule:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movieSchedule *MovieScheduleRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	movieSchedule.model = movieSchedule.model.WithContext(httpContext)

	err := movieSchedule.model.
		Delete(&entities.MovieScheduleEntity{}, id).
		Error
	if err != nil {
		log.Println("Error deleting movie schedule:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movieSchedule *MovieScheduleRepository) IsExistById(httpContext *gin.Context, id uuid.UUID) bool {
	movieSchedule.model = movieSchedule.model.WithContext(httpContext)

	var movieScheduleEntity entities.MovieScheduleEntity
	err := movieSchedule.model.
		First(&movieScheduleEntity, id).
		Error

	if err == gorm.ErrRecordNotFound {
		return false
	} else if err != nil {
		log.Println("Error fetching movie schedule by id", err)
		panic(*exception.ServerErrorException(err))
	}

	return true
}