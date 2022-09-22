package config

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

func JwtParse(token string) (*jwt.Token, jwt.MapClaims, error) {
	tokenParse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != os.Getenv("JWT_SIGNINGKEY") {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(os.Getenv("JWT_KEY")), nil
	})

	if err != nil {
		return nil, nil, err
	}

	claims, ok := tokenParse.Claims.(jwt.MapClaims)
	if !ok || !tokenParse.Valid {
		return nil, nil, errors.New("invalid token")
	}
	return tokenParse, claims, nil
}
