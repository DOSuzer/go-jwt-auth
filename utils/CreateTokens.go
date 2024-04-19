package utils

import (
	"os"
	"time"

	"github.com/DOSuzer/go-jwt-auth/models"
	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("SECRET_KEY")
var refreshSecret = os.Getenv("REFRESH_SECRET_KEY")

var jwtKey = []byte(secret)
var refreshKey = []byte(refreshSecret)

func CreateTokensForUser(user models.User) (string, string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &models.Claims{
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	expirationTimeRefresh := time.Now().Add(24 * 7 * time.Hour)

	claimsRefresh := &models.ClaimsRefresh{
		Role:   user.Role,
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(expirationTimeRefresh),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenRefresh := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)

	t, err := token.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	rt, err := tokenRefresh.SignedString(refreshKey)
	if err != nil {
		return "", "", err
	}

	return t, rt, nil
}
