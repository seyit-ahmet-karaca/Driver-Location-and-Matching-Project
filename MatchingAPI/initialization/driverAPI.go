package initialization

import (
	"MatchingAPI/controller"
	"MatchingAPI/service"
)

// App is a dependency container for the api
type DriverAPI struct {
	Controller DriverController
}

type DriverController struct {
	Driver controller.IMatchingController
}

func InitializeDriverAPI() DriverAPI {
	return DriverAPI{
		Controller:
		DriverController{
				controller.NewMatchingController(
					service.NewMatchingService(),
				),
			},
	}
}
