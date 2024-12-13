package handler

import (
	"net/http"

	domain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/role/dto"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/guard"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/middleware"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"

	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	roleUsecase domain.RoleUsecase
}

func NewRoleHandler(router *gin.Engine, roleUsecase domain.RoleUsecase) {
	roleHandlerRoute := router.Group("/v1/api/roles", guard.AuthGuard())

	roleHandler := &RoleHandler{
		roleUsecase: roleUsecase,
	}

	roleHandlerRoute.DELETE("/", roleHandler.Delete())
	roleHandlerRoute.PATCH("/:id", middleware.ValidateRequestJSON(&dto.RoleUpdateRequest{}), roleHandler.Update())
	roleHandlerRoute.GET("/:id", roleHandler.FindById())
	roleHandlerRoute.GET("/key/:key", roleHandler.Create())
	roleHandlerRoute.POST("/", middleware.ValidateRequestJSON(&dto.RoleCreateRequest{}), roleHandler.Create())
}

func (roleHandler *RoleHandler) Create() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		roleRequest := httpContext.MustGet(constant.RequestBodyJSONKey).(*dto.RoleCreateRequest)
		newRole := domain.RoleEntity{
			Key:  roleRequest.Key,
			Name: roleRequest.Name,
		}
		roleHandler.roleUsecase.Create(&newRole)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (roleHandler *RoleHandler) FindById() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		roleData, _ := roleHandler.roleUsecase.FindById(id)

		httpContext.JSON(http.StatusOK, roleData)
	}
}

func (roleHandler *RoleHandler) FindByKey() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramKey := httpContext.Param("key")
		roleData, _ := roleHandler.roleUsecase.FindByKey(paramKey)

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
		roleHandler.roleUsecase.Update(id, &updateRole)

		httpContext.JSON(http.StatusOK, nil)
	}
}

func (roleHandler *RoleHandler) Delete() gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		paramId := httpContext.Param("id")
		id := utils.UUIDChecker(paramId)
		roleHandler.roleUsecase.Delete(id)

		httpContext.JSON(http.StatusOK, nil)
	}
}
