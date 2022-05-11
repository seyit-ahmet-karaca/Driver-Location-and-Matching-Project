package controller

import (
	_ "driverlocation/docs"
	"driverlocation/model"
	"driverlocation/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

var lock sync.Mutex
var singleInstance IDriverController

type IDriverController interface {
	CreateDriverLocation(e echo.Context) error
	GetNearestDriver(e echo.Context) error
}

type DriverController struct {
	Service service.ILocationService
}

func NewDriverController(service service.ILocationService) IDriverController {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &DriverController{
				Service: service,
			}
		}
	}
	return singleInstance
}

// Create Driver Location
// @Summary Create Driver Location.
// @Description Creates driver location by longitude and latitude or bulk file.
// @Tags Driver
// @Accept */*
// @Produce json
// @Success 200
// @Router /api/driver-locations [post]
// @Param        latitude   query     number  false  "latitude"
// @Param        longitude  query     number  false  "longitude"
// @Param        file       formData  file    false  "bulk location file"
func (c DriverController) CreateDriverLocation(e echo.Context) error {
	file, err := e.FormFile("file")

	if err != nil {
		latitude := e.QueryParam("latitude")
		longitude := e.QueryParam("longitude")
		if latitude == "" || longitude == "" {
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
		point := model.Point{
			Location: model.Location{
				Type:        "Point",
				Coordinates: []float64{longitudeFloat, latitudeFloat},
			},
		}

		err = c.Service.Create(&point)
		if err != nil {
			return e.NoContent(http.StatusInternalServerError)
		}
		return e.NoContent(http.StatusCreated)
	}

	err = c.Service.CreateBulkUpdate(file)
	if err != nil {
		return err
	}

	return e.NoContent(http.StatusCreated)
}

// Get nearest driver at location in radius
// @Summary Gets nearest driver location
// @Description gets nearest driver location by given point and radius
// @Accept  json
// @Produce json
// @Tags Driver
// @Success 200 {object} dto.NearestDriverResponse
// @Router /api/search/location [get]
// @Param        radius     query     int  true  "radius"
// @Param        longitude  query     number  true  "longitude"
// @Param        latitude   query     number  true  "latitude"
func (c DriverController) GetNearestDriver(e echo.Context) error {
	radius := e.QueryParam("radius")
	longitude := e.QueryParam("longitude")
	latitude := e.QueryParam("latitude")
	longitudeFloat, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return err
	}
	latitudeFloat, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return err
	}
	location := &model.Location{
		Type:        "Point",
		Coordinates: []float64{longitudeFloat, latitudeFloat},
	}
	radiusInt, err := strconv.Atoi(radius)
	if err != nil {
		return err
	}
	distance, err := c.Service.GetNearestDriver(location, radiusInt)
	if err != nil {
		return err
	}

	return e.JSON(http.StatusOK, distance)
}
