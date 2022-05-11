package route

import (
	"MatchingAPI/initialization"
	"github.com/labstack/echo/v4"
)

func token(e *echo.Echo, app initialization.UserAPI) {
	e.GET("/token", app.Controller.User.GetJWT)
	e.GET("/tokenNotAuthenticated", app.Controller.User.GetWrongJWT)
}
