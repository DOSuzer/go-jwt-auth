package utils

import (
	"github.com/DOSuzer/go-jwt-auth/models"

	"github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenString string, secret string) (*models.Claims, error) {
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

func ParseRefreshToken(tokenString string, secret string) (*models.ClaimsRefresh, error) {
	key := []byte(secret)

	parsedToken, err := jwt.ParseWithClaims(tokenString, &models.ClaimsRefresh{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, err
	}

	claimsRefresh, ok := parsedToken.Claims.(*models.ClaimsRefresh)

	if !ok {
		return nil, err
	}

	return claimsRefresh, nil
}
