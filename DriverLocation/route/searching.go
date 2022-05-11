package route

import (
	"driverlocation/config"
	"driverlocation/initialization"
	"github.com/labstack/echo/v4"
	"net/http"
)

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		matchApiKey := c.Request().Header.Get(config.Get().MatchApiHeader)
		if matchApiKey != config.Get().SecretMatchApiKey {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "can not request the server",
			})
		}
		return next(c)
	}
}

func searching(g *echo.Group, app initialization.App) {
	searchGroup := g.Group("/search")
	searchGroup.Use(ServerHeader)
	searchGroup.GET("/location", app.Cont.Driver.GetNearestDriver)
}
