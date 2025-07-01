package handler

import (
	"net/http"

	ticketDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/ticket/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/ticket/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	querydto "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/query_dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	ticketUsecase ticketDomain.TicketUsecase
}

func NewTicketHandler(router *gin.Engine, ticketUsecase ticketDomain.TicketUsecase) {
	ticketHandlerRoute := router.Group("/api/v1/tickets", guard.AuthGuard())

	ticketHandler := &TicketHandler{
		ticketUsecase: ticketUsecase,
	}

	// Todo: add permission guard later at second parameter at each route
	ticketHandlerRoute.GET("", ticketHandler.Pagination())
	ticketHandlerRoute.POST("", middleware.ValidateRequestJSON[dto.TicketCreateRequest](), ticketHandler.Create())
	ticketHandlerRoute.GET("/:id", ticketHandler.Detail())
	ticketHandlerRoute.PUT("/:id", middleware.ValidateRequestJSON[dto.TicketUpdateRequest](), ticketHandler.Update())
	ticketHandlerRoute.DELETE("/:id", ticketHandler.Delete())
}

func (ticketHandler *TicketHandler) Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryDto := querydto.AssignFromHttpContext(c)
		data, total := ticketHandler.ticketUsecase.Pagination(c, queryDto)

		formatedData := make([]dto.TicketIndexResponse, len(data))
		for i, ticket := range data {
			formatedData[i] = dto.TicketIndexResponse{
				Id:           ticket.ID,
				Status:       ticket.Status,
				SelectedChairs: ticket.SelectedChairs,
				UpdatedAt:    ticket.UpdatedAt,
			}
		}
		meta := utils.PaginationMetaBuilder(c, total)

		c.JSON(http.StatusOK, utils.SucessResponse(utils.PaginationBuilder(formatedData, *meta)))
	}
}

func (ticketHandler *TicketHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := singleton.GetHTTPRequest[dto.TicketCreateRequest](c)

		ticketHandler.ticketUsecase.Create(c, payload)

		c.JSON(http.StatusCreated, utils.SucessResponse(nil))
	}
}

func (ticketHandler *TicketHandler) Detail() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		ticket := ticketHandler.ticketUsecase.FindOneById(c, id, false)

		c.JSON(http.StatusOK, utils.SucessResponse(dto.NewTicketDetailResponse(*ticket)))
	}
}

func (ticketHandler *TicketHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		payload := singleton.GetHTTPRequest[dto.TicketUpdateRequest](c)

		ticketHandler.ticketUsecase.Update(c, id, payload)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (ticketHandler *TicketHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		ticketHandler.ticketUsecase.Delete(c, id)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}