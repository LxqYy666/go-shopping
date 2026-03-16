package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(key any, method jwt.SigningMethod, claim jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(method, claim)
	return token.SignedString(key)
}

func ParseJWT(key any, jwtStr string) (jwt.Claims, error) {
	token, err := jwt.Parse(jwtStr, func(t *jwt.Token) (any, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("JWT无效")
	}

	return token.Claims, nil
}
