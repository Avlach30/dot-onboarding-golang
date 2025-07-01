package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	movieScheduleDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_schedule/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/ticket/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gorm.io/gorm"
)

type TicketUsecase struct {
	db *gorm.DB
	ticketRepo domain.TicketRepository
	movieScheduleRepo movieScheduleDomain.MovieScheduleRepository
	userRepo userDomain.UserRepository
}

func NewTicketUsecase(db *gorm.DB, ticketRepo domain.TicketRepository, movieScheduleRepo movieScheduleDomain.MovieScheduleRepository, userRepo userDomain.UserRepository) domain.TicketUsecase {
	return &TicketUsecase{
		db: db,
		ticketRepo: ticketRepo,
		movieScheduleRepo: movieScheduleRepo,
		userRepo: userRepo,
	}
}

func (ticketUsecase *TicketUsecase) Pagination(httpContext *gin.Context, queryDto *querydto.QueryDto) ([]entities.TicketEntity, int) {
	return ticketUsecase.ticketRepo.Pagination(httpContext, queryDto)
}

func (ticketUsecase *TicketUsecase) Create(httpContext *gin.Context, payload *entities.TicketEntity) {
	tx := ticketUsecase.db.Begin()

	// Check availability schedule and user
	exception := ticketUsecase.CheckAvailabilityScheduleAndUser(httpContext, payload)
	if exception.StatusCode != 0 {
		tx.Rollback()
		panic(exception)
	}

	// Create ticket
	payload.User = *ticketUsecase.userRepo.FindOneById(httpContext, payload.UserId, false)
	payload.MovieSchedule = *ticketUsecase.movieScheduleRepo.FindOneById(httpContext, payload.MovieScheduleId, false)

	ticketUsecase.ticketRepo.Create(httpContext, payload)

	tx.Commit()
}

func (ticketUsecase *TicketUsecase) Update(httpContext *gin.Context, id uuid.UUID, payload *entities.TicketEntity) {
	tx := ticketUsecase.db.Begin()

	// Check availability schedule and user
	exception := ticketUsecase.CheckAvailabilityScheduleAndUser(httpContext, payload)
	if exception.StatusCode != 0 {
		tx.Rollback()
		panic(exception)
	}

	// Update ticket
	payload.User = *ticketUsecase.userRepo.FindOneById(httpContext, payload.UserId, false)
	payload.MovieSchedule = *ticketUsecase.movieScheduleRepo.FindOneById(httpContext, payload.MovieScheduleId, false)

	ticketUsecase.ticketRepo.Update(httpContext, id, payload)

	tx.Commit()
}

func (ticketUseCase *TicketUsecase) Delete(httpContext *gin.Context, id uuid.UUID) {
	ticketUseCase.ticketRepo.Delete(httpContext, id)
}

func (ticketUsecase *TicketUsecase) FindOneById(httpContext *gin.Context, id uuid.UUID, trashed bool) *entities.TicketEntity {
	return ticketUsecase.ticketRepo.FindOneById(httpContext, id, trashed)
}

func (ticketUsecase *TicketUsecase) CheckAvailabilityScheduleAndUser(httpContext *gin.Context, payload *entities.TicketEntity) exception.Exception {
	err := exception.Exception{}
	movieSchedule := ticketUsecase.movieScheduleRepo.FindOneById(httpContext, payload.MovieScheduleId, false)
	if movieSchedule == nil {
		err = *exception.NotFoundException("Movie Schedule not found")
	}

	user := ticketUsecase.userRepo.FindOneById(httpContext, payload.UserId, false)
	if user == nil {
		err = *exception.NotFoundException("User not found")
	}

	return err
}