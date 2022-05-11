package controller

import (
	"driverlocation/controller"
	"driverlocation/dto"
	"driverlocation/mock"
	"driverlocation/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDriverLocation(t *testing.T) {
	t.Run("Search Driver Location", func(t *testing.T) {
		givenLatitude := 0.1
		givenLongitude := 0.1
		givenRadius := 5

		expectedNearestDriverResponse := &dto.NearestDriverResponse{
			Distance: 10,
			Title:    "Test",
		}

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(
			"/search/location?longitude=%.20f&latitude=%.20f&radius=%d", givenLongitude, givenLatitude, givenRadius),
			http.NoBody)

		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)

		mockController := gomock.NewController(t)
		defer mockController.Finish()

		serviceLocation := &model.Location{
			Type:        "Point",
			Coordinates: []float64{givenLongitude, givenLatitude},
		}

		mockService := mock.NewMockILocationService(mockController)
		mockService.EXPECT().
			GetNearestDriver(serviceLocation, givenRadius).
			Return(expectedNearestDriverResponse, nil).
			Times(1)

		driverController := controller.NewDriverController(mockService)
		result := driverController.GetNearestDriver(context)

		responseDto := &dto.NearestDriverResponse{}
		err := json.Unmarshal(rec.Body.Bytes(), responseDto)

		assert.NoError(t, result)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, responseDto.Distance, expectedNearestDriverResponse.Distance)
		assert.Equal(t, responseDto.Title, expectedNearestDriverResponse.Title)
	})

	t.Run("create driver location with given latitude and longitude", func(t *testing.T) {
		givenLatitude := 0.1
		givenLongitude := 0.1

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(
			"/driver-locations?longitude=%.20f&latitude=%.20f", givenLongitude, givenLatitude),
			http.NoBody)

		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)

		mockController := gomock.NewController(t)
		defer mockController.Finish()

		point := &model.Point{
			Location: model.Location{
				Type:        "Point",
				Coordinates: []float64{givenLongitude, givenLatitude},
			},
		}
		mockService := mock.NewMockILocationService(mockController)
		mockService.EXPECT().
			Create(point).
			Return(nil).
			Times(1)

		driverController := controller.NewDriverController(mockService)
		result := driverController.CreateDriverLocation(context)

		assert.NoError(t, result)
		assert.Equal(t, http.StatusCreated, rec.Code)
	})

	t.Run("error while create driver location with given latitude and longitude", func(t *testing.T) {
		givenLatitude := 0.1
		givenLongitude := 0.1

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(
			"/driver-locations?longitude=%.20f&latitude=%.20f", givenLongitude, givenLatitude),
			http.NoBody)

		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)

		mockController := gomock.NewController(t)
		defer mockController.Finish()

		point := &model.Point{
			Location: model.Location{
				Type:        "Point",
				Coordinates: []float64{givenLongitude, givenLatitude},
			},
		}
		mockService := mock.NewMockILocationService(mockController)
		mockService.EXPECT().
			Create(point).
			Return(errors.New("insert error")).
			Times(1)

		driverController := controller.NewDriverController(mockService)
		result := driverController.CreateDriverLocation(context)

		assert.NoError(t, result)
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
