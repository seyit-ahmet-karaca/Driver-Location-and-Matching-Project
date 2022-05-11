package service

import (
	"driverlocation/exception"
	"driverlocation/mock"
	"driverlocation/model"
	"driverlocation/service"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/umahmood/haversine"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestLocationService(t *testing.T) {
	t.Run("Create single point", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		givenPoint := &model.Point{
			ID:       primitive.ObjectID{},
			Location: model.Location{},
		}

		mockRepository := mock.NewMockILocationRepository(mockController)
		mockRepository.EXPECT().
			Insert(givenPoint).
			Return(nil).
			Times(1)

		locationService := service.NewLocationService(mockRepository)
		err := locationService.Create(givenPoint)

		assert.Nil(t, err)
	})

	t.Run("Get Nearest Driver: successfully", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		givenLocation := &model.Location{
			Type:        "Point",
			Coordinates: []float64{1.0, 1.0},
		}
		givenRadius := 5

		points := []*model.Point{
			&model.Point{
				Title: "testPoint1",
				Location: model.Location{
					Type:        "Point",
					Coordinates: []float64{3.0, 3.0},
				},
			},
			&model.Point{
				Title: "testPoint2",
				Location: model.Location{
					Type:        "Point",
					Coordinates: []float64{5.0, 5.0},
				},
			},
		}

		mockRepository := mock.NewMockILocationRepository(mockController)
		mockRepository.EXPECT().
			FindDriverInDistanceWithRadius(givenLocation, givenRadius).
			Return(points, nil).
			Times(1)

		locationService := service.NewLocationService(mockRepository)
		nearestDriverLocation, err := locationService.GetNearestDriver(givenLocation, givenRadius)

		userCoordinates := haversine.Coord{Lat: givenLocation.Coordinates[1], Lon: givenLocation.Coordinates[0]}
		nearestCoordinates := haversine.Coord{Lat: points[0].Location.Coordinates[1], Lon: points[0].Location.Coordinates[0]}
		_, distance := haversine.Distance(userCoordinates, nearestCoordinates)
		assert.Nil(t, err)
		assert.Equal(t, distance, nearestDriverLocation.Distance)

	})

	t.Run("Get Nearest Driver: successfully", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		givenLocation := &model.Location{
			Type:        "Point",
			Coordinates: []float64{1.0, 1.0},
		}
		givenRadius := 5

		mockRepository := mock.NewMockILocationRepository(mockController)
		mockRepository.EXPECT().
			FindDriverInDistanceWithRadius(givenLocation, givenRadius).
			Return(nil, errors.New("err")).
			Times(1)

		locationService := service.NewLocationService(mockRepository)
		nearestDriverLocation, err := locationService.GetNearestDriver(givenLocation, givenRadius)

		assert.Nil(t, nearestDriverLocation)
		assert.Equal(t, "err",err.Error())

	})

	t.Run("Get Nearest Driver: successfully", func(t *testing.T) {
		mockController := gomock.NewController(t)
		defer mockController.Finish()

		givenLocation := &model.Location{
			Type:        "Point",
			Coordinates: []float64{1.0, 1.0},
		}
		givenRadius := 5

		points := []*model.Point{}

		mockRepository := mock.NewMockILocationRepository(mockController)
		mockRepository.EXPECT().
			FindDriverInDistanceWithRadius(givenLocation, givenRadius).
			Return(points, nil).
			Times(1)

		locationService := service.NewLocationService(mockRepository)
		nearestDriverLocation, err := locationService.GetNearestDriver(givenLocation, givenRadius)

		assert.Nil(t, nearestDriverLocation)
		assert.Equal(t, exception.GetThereIsNoDriverException(givenRadius).Error(), err.Error())
	})
}
