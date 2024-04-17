package utils

import (
	"os"

	"github.com/DOSuzer/go-jwt-auth/models"

	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("SECRET_KEY")

func ParseToken(tokenString string) (*models.Claims, error) {
	key := []byte(secret)

	parsedToken, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*models.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
