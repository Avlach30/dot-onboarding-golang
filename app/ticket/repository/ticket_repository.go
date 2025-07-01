package repository

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/ticket/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
	"gorm.io/gorm"
)

type TicketRepository struct {
	model *gorm.DB
}

func NewTicketRepository(db *gorm.DB) domain.TicketRepository {
	return &TicketRepository{
		model: db.Model(&entities.TicketEntity{}),
	}
}

func (ticketRepository *TicketRepository) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.TicketEntity, int) {
	query := ticketRepository.model.WithContext(httpContext)
	var tickets []entities.TicketEntity
	var total int64

	// Query filter
	query = ticketRepository.queryFilter(query, queryDto)

	// Query sort
	query = ticketRepository.querySort(query, queryDto)

	err := query.Count(&total).
		Error
	if err != nil {
		log.Println("Error counting tickets", err)
		panic(*exception.ServerErrorException(err))
	}

	err = query.Session(&gorm.Session{}).
		Scopes(utils.Paginate(queryDto)).
		Find(&tickets).
		Preload("MovieSchedule").
		Error
	if err != nil {
		log.Println("Error fetching tickets", err)
		panic(*exception.ServerErrorException(err))
	}

	return tickets, int(total)
}

// funcion filter for pagination
func (ticketRepository *TicketRepository) queryFilter(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	if search := queryDto.Search; search != "" {
		query = query.Where("user_id ILIKE ?", search+"%")
	}

	return query
}

// function sort for pagination
func (ticketRepository *TicketRepository) querySort(query *gorm.DB, queryDto *querydto.QueryDto) *gorm.DB {
	sortableColumns := []string{"status", "updated_at"}

	if sortBy := queryDto.SortBy; sortBy != "" {
		if !utils.Contains(sortableColumns, sortBy) {
			panic(*exception.BussinessException("Invalid sortable column"))
		}

		if order := queryDto.Order; order != "" {
			if order != "asc" && order != "desc" {
				panic(*exception.BussinessException("Invalid order"))
			}

			query = query.Order(sortBy + " " + order)
		}
	} else {
		query = query.Order("updated_at desc")
	}

	return query
}

func (ticketRepository *TicketRepository) Create(httpContext *gin.Context, payload *entities.TicketEntity) {
	ticketRepository.model = ticketRepository.model.WithContext(httpContext)

	err := ticketRepository.model.Create(payload).Error
	if err != nil {
		log.Println("Error creating ticket", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (ticketRepository *TicketRepository) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.TicketEntity) {
	ticketRepository.model = ticketRepository.model.WithContext(httpContext)

	err := ticketRepository.model.Where("id = ?", id).Updates(payload).Error
	if err != nil {
		log.Println("Error updating ticket", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (ticketRepository *TicketRepository) Delete(httpContext *gin.Context, id uuid.UUID) {
	ticketRepository.model = ticketRepository.model.WithContext(httpContext)

	err := ticketRepository.model.Delete(&entities.TicketEntity{}, id).Error
	if err != nil {
		log.Println("Error deleting ticket", err)
		panic(*exception.ServerErrorException(err))
	}
}

func (ticketRepository *TicketRepository) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.TicketEntity {
	ticketRepository.model = ticketRepository.model.WithContext(httpContext)
	ticketEntity := &entities.TicketEntity{}
	if trashed {
		ticketRepository.model = ticketRepository.model.Unscoped()
	}

	err := ticketRepository.model.
		Preload("MovieSchedule").
		Preload("MovieSchedule.Movie").
		Preload("MovieSchedule.MovieStudio").
		Preload("User").
		First(&ticketEntity, id).
		Error
	if err == gorm.ErrRecordNotFound {
		panic(*exception.NotFoundException("Ticket not found"))
	} else if err != nil {
		log.Println("Error fetching ticket by id", err)
		panic(*exception.ServerErrorException(err))
	}

	return ticketEntity
}