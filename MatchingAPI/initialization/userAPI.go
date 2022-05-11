package initialization

import (
	"MatchingAPI/controller"
	"MatchingAPI/service"
)

type UserAPI struct {
	Controller UserController
}

type UserController struct {
	User controller.IUserController
}

func InitializeUserAPI() UserAPI {
	return UserAPI{
		Controller:
		UserController{
			controller.NewUserController(
				service.NewUserService(),
			),
		},
	}
}
