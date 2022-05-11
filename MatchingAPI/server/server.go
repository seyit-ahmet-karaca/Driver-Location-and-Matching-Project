package server

import (
	"MatchingAPI/exception"
	"MatchingAPI/route"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func init() {
	//hystrix.ConfigureCommand(route.GetRoute(route.SEARCHING), hystrix.CommandConfig{
	hystrix.ConfigureCommand("/search/driver", hystrix.CommandConfig{
		Timeout:                500,
		RequestVolumeThreshold: 5,
		ErrorPercentThreshold:  50,
		SleepWindow:            hystrix.DefaultSleepWindow,
	})
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if strings.Contains(err.Error(), exception.GetHystrixException().Message) {
		code = http.StatusInternalServerError
		c.JSON(code, exception.GetHystrixException())
		return
	}
	if strings.Contains(err.Error(), exception.GetDataNotFoundException().Message) {
		code = http.StatusNotFound
		c.JSON(code, exception.GetDataNotFoundException())
		return
	}
	if strings.Contains(err.Error(), exception.GetUnauthorizedException().Message) {
		code = http.StatusUnauthorized
		c.JSON(code, exception.GetUnauthorizedException())
		return
	}

	c.JSON(code, exception.GetGenericException(err.Error()))
}

func (server *Server) StartServer(port int) {
	e := echo.New()
	route.Route(e)
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
