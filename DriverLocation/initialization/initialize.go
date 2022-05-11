package initialization

import (
	"driverlocation/controller"
	"driverlocation/repository"
	"driverlocation/service"
)

type App struct {
	Cont Controller
}

type Controller struct {
	Driver controller.IDriverController
}

func InitializeApp() App {
	return App{
		Cont:
			Controller{
				controller.NewDriverController(
					service.NewLocationService(repository.NewLocationRepository()),
				),
			},
	}
}

