package usecase

import (
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	roleRepo domain.AuthRepository
}

// SignIn implements domain.AuthUsecase.
func (authUseCase *AuthUsecase) SignInJWT(email string, password string) string {
	user, err := authUseCase.roleRepo.FindUserByEmail(email)
	if err != nil {
		panic("User and Password not match")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		panic("User and Password not match")
	}

	// Generate JWT token
	authInformation := &domain.AuthEntity{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}

	tokenString, err := jwt.CreateToken(authInformation)
	if err != nil {
		panic("Failed Sign In")
	}

	// Return the token
	return tokenString
}

func NewAuthUsecase(roleRepo domain.AuthRepository) domain.AuthUsecase {
	return &AuthUsecase{
		roleRepo: roleRepo,
	}
}
