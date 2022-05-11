package server

import (
	_ "driverlocation/docs"
	"driverlocation/exception"
	"driverlocation/initialization"
	"driverlocation/route"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	c.JSON(code, exception.GetGenericException(err.Error()))
}

func (server *Server) StartServer(port int) {
	e := echo.New()
	app := initialization.InitializeApp()
	route.Route(e, app)
	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
