package usecase

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	entities "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	userEntities "gitlab.dot.co.id/playground/boilerplates/golang-service/entities"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/sso/ldap"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/sso/oidc"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	authRepo domain.AuthRepository
}

func NewAuthUsecase(authRepo domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{
		authRepo: authRepo,
	}
}

// SignInURLOpenIDClient implements entities.AuthUsecase.
func (authUseCase *AuthUsecase) SignInByOIDCCode(httpContext *gin.Context, code string, redirectUri string) (token string, expiredAt time.Time) {
	// get email from token
	emailSSO, err := oidc.GetEmailByCode(code, redirectUri)
	if err != nil || emailSSO == "" {
		panic(*exception.UnauthorizedException("Not Valid Code"))
	}

	user, err := authUseCase.authRepo.FindUserByEmailWithRoles(httpContext, emailSSO)
	if err != nil || user.ID == uuid.Nil {
		panic(*exception.BussinessException("Email Not Found"))
	}

	SetPermissions(user)
	// sign in jwt
	return authUseCase.CreateJWTToken(user)
}

// SignIn implements entities.AuthUsecase.
func (authUseCase *AuthUsecase) SignInBasic(httpContext *gin.Context, email string, password string) (token string, expirationTime time.Time) {
	user, err := authUseCase.authRepo.FindUserByEmailWithRoles(httpContext, email)
	if err != nil {
		panic(*exception.BussinessException("Email and Password did not match"))
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		panic(*exception.BussinessException("Email and Password did not match"))
	}

	return authUseCase.CreateJWTToken(user)
}

// SignIn implements entities.AuthUsecase.
func (authUseCase *AuthUsecase) SignInLDAP(httpContext *gin.Context, username string, password string) (token string, expirationTime time.Time) {
	userLDAP, err := ldap.AuthUsingLDAP(username, password)
	if err != nil {
		panic(*exception.UnauthorizedException("Email Not Found"))
	}

	user, err := authUseCase.authRepo.FindUserByEmailWithRoles(httpContext, userLDAP.Email)
	if err != nil {
		panic(*exception.UnauthorizedException("Email Not Found"))
	}

	SetPermissions(user)

	return authUseCase.CreateJWTToken(user)
}

// SignIn implements entities.AuthUsecase.
func (authUseCase *AuthUsecase) CreateJWTToken(user *userEntities.UserEntity) (token string, expirationTime time.Time) {

	// Generate JWT token
	authInformation := &entities.AuthEntity{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	expiredInMinutes := config.JwtExpiredInMinutes
	expInMinutes, err := strconv.Atoi(expiredInMinutes)
	if err != nil {
		panic(*exception.ServerErrorException(err))
	}

	expirationDuration := time.Duration(expInMinutes) * time.Minute
	expirationTime = time.Now().Add(expirationDuration)

	tokenString, err := jwt.CreateToken(authInformation, expirationTime)
	if err != nil {
		panic(*exception.ServerErrorException(err))
	}

	SetPermissions(user)

	// Return the token
	return tokenString, expirationTime
}

func SetPermissions(user *userEntities.UserEntity) {
	// set up permissions
	roles := user.Roles
	authPermissions := make([]entities.AuthPermissionEntity, 0)
	for _, role := range roles {
		for _, permission := range role.Permissions {
			authPermissions = append(authPermissions, entities.AuthPermissionEntity{
				ID:   permission.ID,
				Name: permission.Name,
				Key:  permission.Key,
			})
		}
	}

	// set authPermissions in global state
	globalState := singleton.GetGlobalState()
	err := globalState.Set(GenerateHttpContextPermissionKey(user.ID), authPermissions)
	if err != nil {
		panic(*exception.ServerErrorException(err))
	}
}

func GenerateHttpContextPermissionKey(userID uuid.UUID) string {
	return constant.GlobalStatePermissionPrefixKey + userID.String()
}
