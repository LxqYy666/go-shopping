package utils

import (
	"fmt"
	"go-shopping/config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(method jwt.SigningMethod, expiration int, userID int) (string, error) {
	token := jwt.NewWithClaims(method, jwt.MapClaims{
		"exp": expiration,
		"sub": userID,
	})
	return token.SignedString([]byte(config.JWTSecret))
}

func ParseJWT(jwtStr string) (jwt.Claims, error) {
	token, err := jwt.Parse(jwtStr, func(t *jwt.Token) (any, error) {
		return []byte(config.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("JWT无效")
	}

	return token.Claims, nil
}
