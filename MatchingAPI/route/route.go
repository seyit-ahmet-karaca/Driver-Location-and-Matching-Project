package route

import (
	"MatchingAPI/initialization"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//const SEARCHING = "searching"
//
//var route = map[string]string{}
//
//func init() {
//	route[SEARCHING] = "/search/driver"
//}

//func GetRoute(index string) string {
//	return route[index]
//}

func Route(e *echo.Echo) {
	driverAPI := initialization.InitializeDriverAPI()
	searching(e, driverAPI)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	userAPI := initialization.InitializeUserAPI()
	token(e, userAPI)

}
