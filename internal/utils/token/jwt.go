package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func NewToken(secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodPS512)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * 7 * 24)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
