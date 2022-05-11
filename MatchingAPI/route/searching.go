package route

import (
	"MatchingAPI/initialization"
	"MatchingAPI/security"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func searching(e *echo.Echo, app initialization.DriverAPI) {
	jwtConfig := security.GetJwtConfigMiddleware()

	searchGroup := e.Group("/search")
	searchGroup.Use(middleware.JWTWithConfig(jwtConfig))

	searchGroup.GET("/driver", app.Controller.Driver.FindDriver)
	searchGroup.Use()

}
