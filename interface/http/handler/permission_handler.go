package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionHandler struct {
	permissionUsecase domain.PermissionUsecase
}

func NewPermissionHandler(router *gin.Engine, permissionUsecase domain.PermissionUsecase) {
	permissionHandlerRoute := router.Group("/v1/api/permissions")

	permissionHandler := &PermissionHandler{
		permissionUsecase: permissionUsecase,
	}

	permissionHandlerRoute.DELETE("/", permissionHandler.Delete())
	permissionHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.PermissionUpdateRequest{}), permissionHandler.Update())
	permissionHandlerRoute.GET("/:id", permissionHandler.FindById())
	permissionHandlerRoute.GET("/key/:key", permissionHandler.Create())
	permissionHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.PermissionCreateRequest{}), permissionHandler.Create())
}

func (permissionHandler *PermissionHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		permissionRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.PermissionCreateRequest)
		newPermission := domain.PermissionEntity{
			Key:  permissionRequest.Key,
			Name: permissionRequest.Name,
		}
		permissionHandler.permissionUsecase.Create(&newPermission)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (permissionHandler *PermissionHandler) FindById() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id, err := uuid.Parse(paramId)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		permissionData, _ := permissionHandler.permissionUsecase.FindById(id)

		httpContext.JSON(http.StatusOK, permissionData)
	}
}

func (permissionHandler *PermissionHandler) FindByKey() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramKey := httpContext.Param("key")
		permissionData, _ := permissionHandler.permissionUsecase.FindByKey(paramKey)

		httpContext.JSON(http.StatusOK, permissionData)
	}
}

func (permissionHandler *PermissionHandler) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		permissionRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.PermissionUpdateRequest)
		paramId := httpContext.Param("id")
		id, err := uuid.Parse(paramId)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		updatePermission := domain.PermissionEntity{
			Key:  permissionRequest.Key,
			Name: permissionRequest.Name,
		}
		permissionHandler.permissionUsecase.Update(id, &updatePermission)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (permissionHandler *PermissionHandler) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id, err := uuid.Parse(paramId)
		if err != nil {
			httpContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
			return
		}
		permissionHandler.permissionUsecase.Delete(id)

		httpContext.JSON(http.StatusOK, nil)
	}
}
