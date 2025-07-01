package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_studio/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type MovieStudioRepository struct {
	model *gorm.DB
}

func NewMovieStudioRepository(db *gorm.DB) domain.MovieStudioRepository {
	return &MovieStudioRepository{
		model: db.Model(&entities.MovieStudioEntity{}),
	}
}

func (movieStudio *MovieStudioRepository) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.MovieStudioEntity, int) {
	query := movieStudio.model.WithContext(httpContext)
	var movieStudios []entities.MovieStudioEntity
	var total int64

	// Query filter
	query = movieStudio.queryFilter(query, queryDto)
	// Query sort
	query = movieStudio.querySort(query, queryDto)

	err := query.Count(&total).Error
	if err != nil {
		log.Println("Error counting movie studios:", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(queryDto)).
		Find(&movieStudios).Error

	if err != nil {
		log.Println("Error fetching movie studios:", err)
		panic(*exception.ServerErrorException(err))
	}

	return movieStudios, int(total)
}

// Func filter for pagination
func (movieStudio *MovieStudioRepository) queryFilter(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	if search := queryDto.Search; search != "" {
		query = query.Where("name ILIKE ?", search+"%")
	}

	return query
}

func (movieStudio *MovieStudioRepository) querySort(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	sortableColumns := []string{"name", "chair_capacity", "updated_at"}

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

func (movieStudio *MovieStudioRepository) Create(httpContext *gin.Context, payload *entities.MovieStudioEntity) {
	movieStudio.model = movieStudio.model.WithContext(httpContext)

	err := movieStudio.model.Create(payload).Error
	if err != nil {
		log.Println("Error creating movie studio:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movieStudio *MovieStudioRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.MovieStudioEntity {
	movieStudio.model = movieStudio.model.WithContext(httpContext)
	movieStudioEntity := &entities.MovieStudioEntity{}
	if trashed {
		movieStudio.model = movieStudio.model.Unscoped()
	}

	err := movieStudio.model.
		First(&movieStudioEntity, id).
		Error
	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Movie Studio not found"))
	} else if err != nil {
		log.Println("Error movie studio find by id:", err)
		panic(*exception.ServerErrorException(err))
	}

	return movieStudioEntity
}

func (movieStudio *MovieStudioRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.MovieStudioEntity) {
	movieStudio.model = movieStudio.model.WithContext(httpContext)

	err := movieStudio.model.
		Where("id = ?", id).
		Updates(payload).
		Error
	if err != nil {
		log.Println("Error updating movie studio:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movieStudio *MovieStudioRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	movieStudio.model = movieStudio.model.WithContext(httpContext)

	err := movieStudio.model.
		Delete(&entities.MovieStudioEntity{}, id).
		Error
	if err != nil {
		log.Println("Error deleting movie studio:", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (movieStudio *MovieStudioRepository) IsExistsByName(httpContext *gin.Context, name string, id *uuid.UUID) bool {
	movieStudio.model = movieStudio.model.WithContext(httpContext)

	var count int64

	query := movieStudio.model.
		Where("name = ?", name)

	if id != nil {
		query = query.
			Where("id != ?", *id)
	}
	
	err := query.
		Count(&count).
		Error
	if err != nil {
		log.Println("Error checking movie studio exists:", err)
		panic(*exception.ServerErrorException(err))
	}

	return count > 0
}