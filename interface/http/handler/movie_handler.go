package handler

import (
	"log"
	"mime/multipart"
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/master_data/movie/dto"
	fileDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/domain"
	fileDto "gitlab.dot.co.id/playground/boilerplates/golang-service/app/storage/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MovieHandler struct {
	movieUsecase domain.MovieUsecase
	fileUsecase fileDomain.FileUsecase
}

func NewMovieHandler(router *gin.Engine, movieUsecase domain.MovieUsecase, fileUsecase fileDomain.FileUsecase) {
	movieHandlerRoute := router.Group("/api/v1/master-data/movies", guard.AuthGuard())

	movieHandler := &MovieHandler{
		movieUsecase: movieUsecase,
		fileUsecase:  fileUsecase,
	}
	
	// Todo: add permission guard later at second parameter at each route
	movieHandlerRoute.GET("", movieHandler.Pagination())
	movieHandlerRoute.POST("", middleware.ValidateRequestFormData[dto.MovieCreateRequest](), movieHandler.Create())
	movieHandlerRoute.GET("/:id", movieHandler.Detail())
	movieHandlerRoute.PUT("/:id", middleware.ValidateRequestFormData[dto.MovieUpdateRequest](), movieHandler.Update())
	movieHandlerRoute.DELETE("/:id", movieHandler.Delete())
}

func (movieHandler *MovieHandler) Pagination() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, total := movieHandler.movieUsecase.Pagination(c)

		formatedData := make([]dto.MovieIndexResponse, len(data))
		for i, movie := range data {
			formatedData[i] = dto.MovieIndexResponse{
				Id:                movie.ID,
				Title:             movie.Title,
				Genre:             movie.Genre,
				DurationInMinutes: movie.DurationInMinutes,
				UpdatedAt:         movie.UpdatedAt,
			}
		}
		meta := utils.PaginationMetaBuilder(c, total)

		c.JSON(http.StatusOK, utils.SucessResponse(utils.PaginationBuilder(formatedData, *meta)))
	}
}

func (movieHandler *MovieHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		payload := singleton.GetHTTPRequest[dto.MovieCreateRequest](c)

		// Upload poster
		posterUrl := movieHandler.uploadPoster(c, payload.Poster)
		
		newMovie := entities.MovieEntity{
			Title:             payload.Title,
			Genre:             payload.Genre,
			DurationInMinutes: payload.DurationInMinutes,
			PosterUrl:         posterUrl,
			Desciption:        payload.Description,
		}

		movieHandler.movieUsecase.Create(c, &newMovie)

		c.JSON(http.StatusCreated, utils.SucessResponse(nil))
	}
}

func (movieHandler *MovieHandler) Detail() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		movie := movieHandler.movieUsecase.FindOneById(c, id, false)

		c.JSON(http.StatusOK, utils.SucessResponse(dto.NewMovieDetailResponse(*movie)))
	}
}

func (movieHandler *MovieHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		payload := singleton.GetHTTPRequest[dto.MovieUpdateRequest](c)

		// Delete old poster
		err := movieHandler.DeleteExistingPosterUrl(c, id);
		if (err != nil){
			log.Println("Error delete old poster", err)
			panic(*exception.ServerErrorException(err))
		}

		// Upload poster
		posterUrl := movieHandler.uploadPoster(c, payload.Poster)
		
		movie := entities.MovieEntity{
			Title:             payload.Title,
			Genre:             payload.Genre,
			DurationInMinutes: payload.DurationInMinutes,
			PosterUrl:         posterUrl,
			Desciption:        payload.Description,
		}

		movieHandler.movieUsecase.Update(c, id, &movie)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (movieHandler *MovieHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		paramId := c.Param("id")
		id := utils.UUIDChecker(paramId)

		movieHandler.movieUsecase.Delete(c, id)

		c.JSON(http.StatusOK, utils.SucessResponse(nil))
	}
}

func (movieHandler *MovieHandler) DeleteExistingPosterUrl(c *gin.Context, id uuid.UUID) error {
	// Get movie by id
	movieData := movieHandler.movieUsecase.FindOneById(c, id, false)
	existingPosterUrl := movieData.PosterUrl
	// Delete old poster
	err := movieHandler.fileUsecase.Delete(c, []string{existingPosterUrl})

	return err
}

func (movieHandler *MovieHandler) uploadPoster(c *gin.Context, file multipart.FileHeader) string {
	uploadReq := &fileDto.UploadFilesRequest{
		Files: []multipart.FileHeader{file},
	}

	uploadedPaths, err := movieHandler.fileUsecase.UploadFiles(c, uploadReq)
	if err != nil {
		log.Println("Error uploading files", err)
		panic(*exception.ServerErrorException(err))
	}
	fileUrl := uploadedPaths[file.Filename]
	return fileUrl
}