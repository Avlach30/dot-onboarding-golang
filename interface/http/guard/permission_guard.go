package guard

import (
	"github.com/gin-gonic/gin"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/usecase"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
	state "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/utils"
)

// Custom guard to check JWT token in Authorization header
func PermissionGuard(permissionsToCheck ...string) gin.HandlerFunc {
	return func(httpContext *gin.Context) {
		isAuthorized, isExists := httpContext.Get(constant.IsAuthorizedHeaderKey)
		if !isExists || !isAuthorized.(bool) {
			panic(*exception.UnauthorizedException("Not Authorized"))
		}

		authEntityInfo, isExists := httpContext.Get(constant.AuthUserInfoKey)
		if !isExists || !isAuthorized.(bool) {
			panic(*exception.UnauthorizedException("Not Authorized"))
		}

		userID := authEntityInfo.(*jwt.CustomClaims).ID

		// get global state
		globalState := state.GetGlobalState()
		permissions := state.Get[[]domain.AuthPermission](usecase.GenerateHttpContextPermissionKey(userID), globalState)

		if permissions == nil {
			panic(*exception.ForbiddenException("Not Allowed Access"))
		}

		permissionInState := permissions.([]domain.AuthPermission)
		permissionKeys := authPermissionMapToKeys(&permissionInState)

		// check user permission is allowed access or not
		isPermissionsIntersect := utils.IsAnyIntersect(permissionsToCheck, permissionKeys)
		if isPermissionsIntersect {
			panic(*exception.ForbiddenException("Not Allowed Access"))
		}

		// set in context that permission is checked
		httpContext.Set(constant.IsPemissionCheckedKey, true)

		httpContext.Next()
	}
}

func authPermissionMapToKeys(authPermissionEntities *[]domain.AuthPermission) []string {
	entities := *authPermissionEntities
	keys := make([]string, len(entities))
	for i, entity := range entities {
		keys[i] = entity.Key
	}

	return keys
}
