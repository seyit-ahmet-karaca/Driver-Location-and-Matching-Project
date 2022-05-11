package service

import (
	"MatchingAPI/config"
	"MatchingAPI/dto"
	"github.com/golang-jwt/jwt"
	"sync"
	"time"
)

var userLock = &sync.Mutex{}
var userSingleInstance IUserService

type IUserService interface {
	CreateJWT() *dto.UserResponse
	CreateWrongJWT() *dto.UserResponse
}

type UserService struct{}

func NewUserService() IUserService {
	if userSingleInstance == nil {
		userLock.Lock()
		defer userLock.Unlock()
		if userSingleInstance == nil {
			userSingleInstance = &UserService{}
		}
	}
	return userSingleInstance
}

func (u *UserService) CreateJWT() *dto.UserResponse {
	claims := &dto.JwtCustomClaims{
		Name:          "Jon Snow",
		Authenticated: true,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.Get().JwtDurationHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(config.Get().JwtSecret))

	return &dto.UserResponse{
		Name:  claims.Name,
		Token: "Bearer " + signedToken,
	}
}

func (u *UserService) CreateWrongJWT() *dto.UserResponse {
	claims := &dto.JwtCustomClaims{
		Name:          "Joffrey Baratheon",
		Authenticated: false,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(config.Get().JwtDurationHour)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(config.Get().JwtSecret))

	return &dto.UserResponse{
		Name:  claims.Name,
		Token: "Bearer " + signedToken,
	}
}
