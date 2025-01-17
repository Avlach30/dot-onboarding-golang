package usecase

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
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

// SignInURLOpenIDClient implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) SignInByOIDCCode(httpContext *gin.Context, code string) (token string, expiredAt time.Time) {

	// get email from token
	emailSSO, err := oidc.GetEmailByCode(code)
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

// SignIn implements domain.AuthUsecase.
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

// SignIn implements domain.AuthUsecase.
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

// SignIn implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) CreateJWTToken(user *userDomain.UserEntity) (token string, expirationTime time.Time) {

	// Generate JWT token
	authInformation := &domain.AuthEntity{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	expiredInDays := config.JwtExpiredInDays
	expInDays, err := strconv.Atoi(expiredInDays)
	if err != nil {
		panic(*exception.ServerErrorException(err))
	}

	expirationDuration := time.Duration(expInDays) * 24 * time.Hour
	expirationTime = time.Now().Add(expirationDuration)

	tokenString, err := jwt.CreateToken(authInformation, expirationTime)
	if err != nil {
		panic(*exception.ServerErrorException(err))
	}

	SetPermissions(user)

	// Return the token
	return tokenString, expirationTime
}

func NewAuthUsecase(authRepo domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{
		authRepo: authRepo,
	}
}

func SetPermissions(user *userDomain.UserEntity) {
	// set up permissions
	roles := user.Roles
	authPermissions := make([]domain.AuthPermissionEntity, 0)
	for _, role := range roles {
		for _, permission := range role.Permissions {
			authPermissions = append(authPermissions, domain.AuthPermissionEntity{
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
