package handler

import (
	"fmt"
	"net/http"

	movieDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie/domain"
	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_schedule/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_schedule/dto"
	movieStudioDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_studio/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieScheduleHandler struct {
	movieScheduleUsecase domain.MovieScheduleUsecase
	movieStudioUsecase movieStudioDomain.MovieStudioUsecase
	movieUsecase movieDomain.MovieUsecase
}

func NewMovieScheduleHandler(router *gin.Engine, movieScheduleUsecase domain.MovieScheduleUsecase, movieStudioUsecase movieStudioDomain.MovieStudioUsecase, movieUsecase movieDomain.MovieUsecase) {
	movieScheduleHandlerRoute := router.Group("/api/v1/master-data/movie-schedules", guard.AuthGuard())

	movieScheduleHandler := &MovieScheduleHandler{
		movieScheduleUsecase: movieScheduleUsecase,
		movieStudioUsecase: movieStudioUsecase,
		movieUsecase: movieUsecase,
	}
	
	// Todo: add permission guard later at second parameter at each route
	movieScheduleHandlerRoute.GET("", movieScheduleHandler.Pagination())
	movieScheduleHandlerRoute.POST("", middleware.ValidateRequestJSON[dto.MovieScheduleCreateRequest](), movieScheduleHandler.Create())
	movieScheduleHandlerRoute.GET("/:id", movieScheduleHandler.Detail())
	movieScheduleHandlerRoute.PUT("/:id", middleware.ValidateRequestJSON[dto.MovieScheduleUpdateRequest](), movieScheduleHandler.Update())
	movieScheduleHandlerRoute.DELETE("/:id", movieScheduleHandler.Delete())
}

func (movieScheduleHandler *MovieScheduleHandler) Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, total := movieScheduleHandler.movieScheduleUsecase.Pagination(c)

		formatedData := make([]dto.MovieScheduleIndexResponse, len(data))
		for i, movieSchedule := range data {
			formatedData[i] = dto.MovieScheduleIndexResponse{
				Id:           movieSchedule.ID,
				ShowDatetime: movieSchedule.ShowDatetime,
				Price:        movieSchedule.Price,
				UpdatedAt:    movieSchedule.UpdatedAt,
			}
		}
		meta := utils.PaginationMetaBuilder(c, total)

		c.JSON(http.StatusOK, utils.SucessResponse(utils.PaginationBuilder(formatedData, *meta)))
	}
}

func (movieScheduleHandler *MovieScheduleHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := singleton.GetHTTPRequest[dto.MovieScheduleCreateRequest](c)

		// Check if movie and movie studio exists
		movie := movieScheduleHandler.CheckIsExistsMovieById(c, payload.MovieID)
		movieStudio :=movieScheduleHandler.CheckIsExistsMovieStudioById(c, payload.MovieStudioID)

		fmt.Println(movie)
		fmt.Println(*movie)
		
		newMovieSchedule := entities.MovieScheduleEntity{
			Movie:         *movie,
			MovieStudio:   *movieStudio,
			ShowDatetime: payload.ShowDatetime,
			Price:        payload.Price,
		}

		movieScheduleHandler.movieScheduleUsecase.Create(c, &newMovieSchedule)

		c.JSON(http.StatusCreated, utils.SucessResponse(nil))
	}
}

func (movieScheduleHandler *MovieScheduleHandler) Detail() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		movieSchedule := movieScheduleHandler.movieScheduleUsecase.FindOneById(c, id, false)

		c.JSON(http.StatusOK, utils.SucessResponse(dto.NewMovieScheduleDetailResponse(*movieSchedule)))
	}
}

func (movieScheduleHandler *MovieScheduleHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		payload := singleton.GetHTTPRequest[dto.MovieScheduleUpdateRequest](c)

		// Check if movie and movie studio exists
		movie := movieScheduleHandler.CheckIsExistsMovieById(c, payload.MovieID)
		movieStudio := movieScheduleHandler.CheckIsExistsMovieStudioById(c, payload.MovieStudioID)
		
		movieSchedule := entities.MovieScheduleEntity{
			Movie:         *movie,
			MovieStudio:   *movieStudio,
			ShowDatetime: payload.ShowDatetime,
			Price:        payload.Price,
		}

		movieScheduleHandler.movieScheduleUsecase.Update(c, id, &movieSchedule)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (movieScheduleHandler *MovieScheduleHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		movieScheduleHandler.movieScheduleUsecase.Delete(c, id)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (movieScheduleHandler *MovieScheduleHandler) CheckIsExistsMovieById(c *gin.Context, id uuid.UUID) *entities.MovieEntity {
	movieData := movieScheduleHandler.movieUsecase.FindOneById(c, id, false)
	if movieData == nil {
		c.JSON(http.StatusNotFound, utils.SucessResponse(nil))
	}
	return movieData
}

func (movieScheduleHandler *MovieScheduleHandler) CheckIsExistsMovieStudioById(c *gin.Context, id uuid.UUID) *entities.MovieStudioEntity {
	movieStudioData := movieScheduleHandler.movieStudioUsecase.FindOneById(c, id, false)
	if movieStudioData == nil {
		c.JSON(http.StatusNotFound, utils.SucessResponse(nil))
	}
	return movieStudioData
}