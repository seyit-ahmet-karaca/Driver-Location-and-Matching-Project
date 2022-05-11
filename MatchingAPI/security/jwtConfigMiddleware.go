package security

import (
	"MatchingAPI/config"
	"MatchingAPI/dto"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"strings"
)

func GetJwtConfigMiddleware() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &dto.JwtCustomClaims{},
		SigningKey: []byte(config.Get().JwtSecret),
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			return TokenParser(auth)
		},
	}
}

func TokenParser(auth string) (interface{}, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}
		return []byte(config.Get().JwtSecret), nil
	}
	if strings.HasPrefix(auth, "Bearer ") {
		auth = strings.Replace(auth,"Bearer ","",1)
	}
	token, err := jwt.Parse(auth, keyFunc)
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("invalid token")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		authenticated := claims["authenticated"]
		if isAuthenticated, isOK := authenticated.(bool); isOK {
			if !isAuthenticated {
				return nil, errors.New("the user has not suitable credential")
			}
		}
	}
	return token, nil
}
