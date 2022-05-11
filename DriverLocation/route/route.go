package route

import (
	"driverlocation/initialization"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func Route(e *echo.Echo, app initialization.App) {
	api := e.Group("/api")
	driverLocation(api, app)
	searching(api, app)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
