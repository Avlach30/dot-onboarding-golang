package guard

import (
	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/usecase"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

// Custom guard to check JWT token in Authorization header
func PermissionGuard(permissionsToCheck ...string) gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		isAuthorized, isExists := httpContext.Get(constant.IsAuthorizedHeaderKey)
		if !isExists || !isAuthorized.(bool) {
			panic(*exception.UnauthorizedException("Not Authorized"))
		}

		userID := singleton.GetAuthUserID(httpContext)

		// get global state
		globalState := singleton.GetGlobalState()
		permissions, err := singleton.Get[[]entities.AuthPermissionEntity](usecase.GenerateHttpContextPermissionKey(userID), globalState)
		if err != nil {
			panic(*exception.UnauthorizedException("Try again later"))
		}

		if permissions == nil {
			panic(*exception.ForbiddenException("Not Allowed Access"))
		}

		permissionInState := permissions.([]entities.AuthPermissionEntity)
		permissionKeys := authPermissionEntityMapToKeys(&permissionInState)

		// check user permission is allowed access or not
		isPermissionsIntersect := utils.IsAnyIntersect(permissionsToCheck, permissionKeys)
		if isPermissionsIntersect == false {
			panic(*exception.ForbiddenException("Not Allowed Access"))
		}

		// set in context that permission is checked
		httpContext.Set(constant.IsPemissionCheckedKey, true)

		httpContext.Next()
	}
}

func authPermissionEntityMapToKeys(authPermissionEntities *[]entities.AuthPermissionEntity) []string {
	entities := *authPermissionEntities
	keys := make([]string, len(entities))
	for i, entity := range entities {
		keys[i] = entity.Key
	}

	return keys
}
