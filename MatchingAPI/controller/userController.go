package controller

import (
	_ "MatchingAPI/docs"
	"MatchingAPI/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync"
)

var userLock sync.Mutex
var userSingletonInstance IUserController

type IUserController interface {
	GetJWT(e echo.Context) error
	GetWrongJWT(e echo.Context) error
}

type UserController struct {
	Service service.IUserService
}

func NewUserController(service service.IUserService) IUserController {
	if userSingletonInstance == nil {
		userLock.Lock()
		defer userLock.Unlock()
		if userSingletonInstance == nil {
			userSingletonInstance = &UserController{
				Service: service,
			}
		}
	}
	return userSingletonInstance
}

// Get valid jwt token
// @Summary Get valid jwt token
// @Description Gets valid jwt token to send request matching API
// @Tags JWT
// @Accept */*
// @Produce json
// @Success 200
// @Router /token [get]
func (m *UserController) GetJWT(e echo.Context) error {
	return e.JSON(http.StatusOK, m.Service.CreateJWT())
}

// Get unauthanticated jwt token
// @Summary Get unauthanticated jwt token
// @Description Gets unauthanticated jwt token to send request matching API
// @Accept */*
// @Tags JWT
// @Produce json
// @Success 200
// @Router /tokenNotAuthenticated [get]
func (m *UserController) GetWrongJWT(e echo.Context) error {
	return e.JSON(http.StatusOK, m.Service.CreateWrongJWT())
}
