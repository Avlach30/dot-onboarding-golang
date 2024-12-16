package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/app/auth/domain"
	"gitlab.dot.co.id/playground/boilerplates/golang-service/config"
)

type CustomClaims struct {
	ID    uuid.UUID `json:"ID"`
	Email string    `json:"email"`
	Name  string    `json:"name"`
	jwt.RegisteredClaims
}

// CreateToken generates a new JWT token
func CreateToken(authEntity *domain.AuthEntity, expirationTime time.Time) (res string, err error) {
	secret := []byte(config.Secret)

	claims := CustomClaims{
		ID:    authEntity.ID,
		Email: authEntity.Email,
		Name:  authEntity.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    authEntity.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}

// ParseToken parses a JWT token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
