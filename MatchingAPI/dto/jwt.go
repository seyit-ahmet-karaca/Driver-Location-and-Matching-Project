package dto

import "github.com/golang-jwt/jwt"

type JwtCustomClaims struct {
	Name          string `json:"name"`
	Authenticated bool   `json:"authenticated"`
	jwt.StandardClaims
}

