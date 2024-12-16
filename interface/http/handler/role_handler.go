package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleUsecase domain.RoleUsecase
}

func NewRoleHandler(router *gin.Engine, roleUsecase domain.RoleUsecase) {
	// roleHandlerRoute := router.Group("/v1/api/roles", guard.AuthGuard())
	roleHandlerRoute := router.Group("/v1/api/roles")

	roleHandler := &RoleHandler{
		roleUsecase: roleUsecase,
	}

	roleHandlerRoute.DELETE("/:id", roleHandler.Delete())
	roleHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.RoleUpdateRequest{}), roleHandler.Update())
	roleHandlerRoute.GET("/:id", roleHandler.FindById())
	roleHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.RoleCreateRequest{}), roleHandler.Create())
}

func (roleHandler *RoleHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		roleRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.RoleCreateRequest)
		newRole := domain.RoleEntity{
			Key:  roleRequest.Key,
			Name: roleRequest.Name,
		}

		roleHandler.roleUsecase.Create(singleton.GetContextFromGinContext(httpContext), &newRole)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (roleHandler *RoleHandler) FindById() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		roleData, _ := roleHandler.roleUsecase.FindById(singleton.GetContextFromGinContext(httpContext), id)

		httpContext.JSON(http.StatusOK, roleData)
	}
}

func (roleHandler *RoleHandler) FindByKey() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramKey := httpContext.Param("key")
		roleData, _ := roleHandler.roleUsecase.FindByKey(singleton.GetContextFromGinContext(httpContext), paramKey)

		httpContext.JSON(http.StatusOK, roleData)
	}
}

func (roleHandler *RoleHandler) Update() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		roleRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.RoleUpdateRequest)
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		updateRole := domain.RoleEntity{
			Key:  roleRequest.Key,
			Name: roleRequest.Name,
		}
		roleHandler.roleUsecase.Update(singleton.GetContextFromGinContext(httpContext), id, &updateRole)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (roleHandler *RoleHandler) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		roleHandler.roleUsecase.Delete(singleton.GetContextFromGinContext(httpContext), id)

		httpContext.JSON(http.StatusOK, nil)
	}
}
