package usecase

import (
	"fmt"

	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/interface/http/exception"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
)

type AuthUsecase struct {
	roleRepo domain.AuthRepository
}

// SignIn implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) SignInJWT(email string, password string) string {
	user, err := authUseCase.roleRepo.FindUserByEmail(email)
	if err != nil {
		panic(*exception.BussinessException("Email and Password did not match"))
	}

	fmt.Println(user)

	// if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
	// 	panic(*exception.BussinessException("Email and Password did not match"))
	// }

	// Generate JWT token
	authInformation := &domain.AuthEntity{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	tokenString, err := jwt.CreateToken(authInformation)
	if err != nil {
		panic(*exception.ServerErrorException("Server Error"))
	}

	// Return the token
	return tokenString
}

func NewAuthUsecase(roleRepo domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{
		roleRepo: roleRepo,
	}
}
