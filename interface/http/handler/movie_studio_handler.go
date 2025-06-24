package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_studio/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie_studio/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type MovieStudioHandler struct {
	movieStudioUsecase domain.MovieStudioUsecase
}

func NewMovieStudioHandler(router *gin.Engine, movieStudioUsecase domain.MovieStudioUsecase) {
	movieStudioHandlerRoute := router.Group("/api/v1/master-data/movie-studios", guard.AuthGuard())

	movieStudioHandler := &MovieStudioHandler{
		movieStudioUsecase: movieStudioUsecase,
	}
	
	// Todo: add permission guard later at second parameter at each route
	movieStudioHandlerRoute.GET("", movieStudioHandler.Pagination())
	movieStudioHandlerRoute.POST("", middleware.ValidateRequestJSON[dto.MovieStudioCreateRequest](), movieStudioHandler.Create())
	movieStudioHandlerRoute.GET("/:id", movieStudioHandler.Detail())
	movieStudioHandlerRoute.PUT("/:id", middleware.ValidateRequestJSON[dto.MovieStudioUpdateRequest](), movieStudioHandler.Update())
	movieStudioHandlerRoute.DELETE("/:id", movieStudioHandler.Delete())
}

func (movieStudioHandler *MovieStudioHandler) Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, total := movieStudioHandler.movieStudioUsecase.Pagination(c)

		formatedData := make([]dto.MovieStudioIndexResponse, len(data))
		for i, movieStudio := range data {
			formatedData[i] = dto.MovieStudioIndexResponse{
				Id:            movieStudio.ID,
				Name:          movieStudio.Name,
				ChairCapacity: movieStudio.ChairCapacity,
				CreatedAt:     movieStudio.CreatedAt,
			}
		}
		meta := utils.PaginationMetaBuilder(c, total)

		c.JSON(http.StatusOK, utils.SucessResponse(utils.PaginationBuilder(formatedData, *meta)))
	}
}

func (movieStudioHandler *MovieStudioHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := singleton.GetHTTPRequest[dto.MovieStudioCreateRequest](c)
		
		newMovieStudio := entities.MovieStudioEntity{
			Name:          payload.Name,
			ChairCapacity: payload.ChairCapacity,
			AdditionalCapacities: utils.StringArray(payload.AdditionalCapacities),
		}

		movieStudioHandler.movieStudioUsecase.Create(c, &newMovieStudio)

		c.JSON(http.StatusCreated, utils.SucessResponse(nil))
	}
}

func (movieStudioHandler *MovieStudioHandler) Detail() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		movieStudio := movieStudioHandler.movieStudioUsecase.FindOneById(c, id, false)

		c.JSON(http.StatusOK, utils.SucessResponse(dto.NewMovieStudioDetailResponse(*movieStudio)))
	}
}

func (movieStudioHandler *MovieStudioHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		payload := singleton.GetHTTPRequest[dto.MovieStudioUpdateRequest](c)
		
		movieStudio := entities.MovieStudioEntity{
			Name:          payload.Name,
			ChairCapacity: payload.ChairCapacity,
			AdditionalCapacities: payload.AdditionalCapacities,
		}

		movieStudioHandler.movieStudioUsecase.Update(c, id, &movieStudio)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (movieStudioHandler *MovieStudioHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		movieStudioHandler.movieStudioUsecase.Delete(c, id)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}