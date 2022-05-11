package controller

import (
	_ "MatchingAPI/docs"
	"MatchingAPI/dto"
	"MatchingAPI/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

var lock sync.Mutex
var singleInstance IMatchingController

type IMatchingController interface {
	FindDriver(e echo.Context) error
}

type MatchingController struct {
	Service service.IMatchingService
}

func NewMatchingController(service service.IMatchingService) IMatchingController {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &MatchingController{
				Service: service,
			}
		}
	}
	return singleInstance
}

// Find nearest driver to user
// @Summary Find nearest driver to user
// @Description Find nearest driver to user from given location point
// @Accept */*
// @Tags User-Driver Matching
// @Produce json
// @Success 200 {object} dto.MatchResponse
// @Router /search/driver [get]
// @Param   latitude   		  query     number  true  "latitude"
// @Param   longitude  		  query     number  true  "longitude"
// @Param   radius     		  query     int     true  "radius"
// @Param   Authorization     header    string  true  "Authorization"
// @Security ApiKeyAuth
func (m *MatchingController) FindDriver(e echo.Context) error {
	radius := e.QueryParam("radius")
	latitude := e.QueryParam("latitude")
	longitude := e.QueryParam("longitude")

	radiusInt, err := strconv.Atoi(radius)
	if err != nil {
		return err
	}
	latitudeFloat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return err
	}
	longitudeFloat, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return err
	}
	location := &dto.SearchLocation{
		Radius:    radiusInt,
		Latitude:  latitudeFloat,
		Longitude: longitudeFloat,
	}
	client := &http.Client{}
	if err != nil {
		return err
	}
	driverWithDistance, err := m.Service.FindDriver(location, client)
	if err != nil {
		return err
	}
	return e.JSON(http.StatusOK, driverWithDistance)
}
