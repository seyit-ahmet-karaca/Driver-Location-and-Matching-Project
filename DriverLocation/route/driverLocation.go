package route

import (
	"driverlocation/initialization"
	"github.com/labstack/echo/v4"
)

func driverLocation(g *echo.Group, app initialization.App) {
	driverLocation := g.Group("/driver-locations")
	driverLocation.POST("", app.Cont.Driver.CreateDriverLocation)
}
