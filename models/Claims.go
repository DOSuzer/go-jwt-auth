package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type ClaimsRefresh struct {
	Role   string `json:"role"`
	UserId uint   `json:"user_id"`
	jwt.RegisteredClaims
}
