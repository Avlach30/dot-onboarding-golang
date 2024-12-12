package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(router *gin.Engine, userUsecase domain.UserUsecase) {
	userHandlerRoute := router.Group("/v1/api/users")

	userHandler := &UserHandler{
		userUsecase: userUsecase,
	}

	userHandlerRoute.DELETE("/", userHandler.Delete())
	userHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.UserUpdateRequest{}), userHandler.Update())
	userHandlerRoute.GET("/:id", userHandler.FindById())
	userHandlerRoute.GET("/key/:key", userHandler.Create())
	userHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.UserCreateRequest{}), userHandler.Create())
}

func (userHandler *UserHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		userRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.UserCreateRequest)
		newUser := domain.UserEntity{
			Name:     userRequest.Name,
			Email:    userRequest.Email,
			Password: userRequest.Password,
		}
		userHandler.userUsecase.Create(&newUser)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (userHandler *UserHandler) FindById() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id, err := uuid.Parse(paramId)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		userData, _ := userHandler.userUsecase.FindById(id, false)

		httpContext.JSON(http.StatusOK, userData)
	}
}

func (userHandler *UserHandler) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		userRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.UserUpdateRequest)
		paramId := httpContext.Param("id")
		id, err := uuid.Parse(paramId)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		updateUser := domain.UserEntity{
			Name:  userRequest.Name,
			Email: userRequest.Email,
		}
		userHandler.userUsecase.Update(id, &updateUser)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (userHandler *UserHandler) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id, err := uuid.Parse(paramId)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		userHandler.userUsecase.Delete(id)

		httpContext.JSON(http.StatusOK, nil)
	}
}
