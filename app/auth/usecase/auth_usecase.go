package usecase

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/permission/immutable"
	userDomain "gitlab.dot.co.id/playground/boilerplates/golang-service/app/user/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/constant"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
	state "gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/singleton"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/sso/ldap"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/sso/oidc"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	authRepo domain.AuthRepository
}

// SignInURLOpenIDClient implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) SignInByOIDCCode(context *context.Context, code string) (token string, expiredAt time.Time) {

	// get email from token
	emailSSO, err := oidc.GetEmailByCode(code)
	if err != nil || emailSSO == "" {
		panic(*exception.UnauthorizedException("Not Valid Code"))
	}

	log.Println(emailSSO)

	user, err := authUseCase.authRepo.FindUserByEmail(context, emailSSO)
	if err != nil || user.ID == uuid.Nil {
		panic(*exception.BussinessException("Email Not Found"))
	}

	// sign in jwt
	return authUseCase.CreateJWTToken(user)
}

// SignIn implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) SignInBasic(context *context.Context, email string, password string) (token string, expirationTime time.Time) {
	user, err := authUseCase.authRepo.FindUserByEmail(context, email)
	if err != nil {
		panic(*exception.BussinessException("Email and Password did not match"))
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		panic(*exception.BussinessException("Email and Password did not match"))
	}

	return authUseCase.CreateJWTToken(user)
}

// SignIn implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) SignInLDAP(context *context.Context, username string, password string) (token string, expirationTime time.Time) {
	userLDAP, err := ldap.AuthUsingLDAP(username, password)
	if err != nil {
		panic(*exception.UnauthorizedException("Email Not Found"))
	}

	user, err := authUseCase.authRepo.FindUserByEmail(context, userLDAP.Email)
	if err != nil {
		panic(*exception.UnauthorizedException("Email Not Found"))
	}

	return authUseCase.CreateJWTToken(user)
}

// SignIn implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) CreateJWTToken(user *userDomain.UserEntity) (token string, expirationTime time.Time) {
	// FIX ME: dummy permission
	dummyPermissions := []domain.AuthPermissionEntity{
		{
			ID:   uuid.New(),
			Name: immutable.PermissionApprove,
			Key:  immutable.PermissionApprove,
		},
		{
			ID:   uuid.New(),
			Name: immutable.PermissionExport,
			Key:  immutable.PermissionExport,
		},
	}

	// set permissions in global state
	globalState := state.GetGlobalState()
	globalState.Set(GenerateHttpContextPermissionKey(user.ID), &dummyPermissions)

	// Generate JWT token
	authInformation := &domain.AuthEntity{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	expiredInDays := config.JwtExpiredInDays
	expInDays, err := strconv.Atoi(expiredInDays)
	if err != nil {
		panic(*exception.ServerErrorException("Error Create Token"))
	}

	expirationDuration := time.Duration(expInDays) * 24 * time.Hour
	expirationTime = time.Now().Add(expirationDuration)

	tokenString, err := jwt.CreateToken(authInformation, expirationTime)
	if err != nil {
		panic(*exception.ServerErrorException("Error Create Token"))
	}

	// Return the token
	return tokenString, expirationTime
}

func NewAuthUsecase(authRepo domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{
		authRepo: authRepo,
	}
}

func GenerateHttpContextPermissionKey(userID uuid.UUID) string {
	return constant.GlobalStatePermissionPrefixKey + userID.String()
}
