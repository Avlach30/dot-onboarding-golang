package domain

type AuthUsecase interface {
	SignInJWT(email string, password string) string
}
