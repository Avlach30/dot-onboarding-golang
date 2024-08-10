package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/codespace-id/codespace-x/config"
	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	Role            string `json:"role"`
	PhoneNumber     string `json:"phone_number"`
	FirebaseIdToken string `json:"firebase_id_token"`
	jwt.RegisteredClaims
}

// CreateToken generates a new JWT token
func CreateToken(phoneNumber, role, firebaseIdToken string) (res string, err error) {
	expiredInDays := config.JwtExpiredInDays
	secret := []byte(config.Secret)

	expInDays, err := strconv.Atoi(expiredInDays)
	if err != nil {
		return res, errors.New("CreateToken.Strconv.AtoiFailed")
	}

	expirationDuration := time.Duration(expInDays) * 24 * time.Hour
	expirationTime := time.Now().Add(expirationDuration)

	claims := CustomClaims{
		PhoneNumber:     phoneNumber,
		Role:            role,
		FirebaseIdToken: firebaseIdToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    phoneNumber,
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
