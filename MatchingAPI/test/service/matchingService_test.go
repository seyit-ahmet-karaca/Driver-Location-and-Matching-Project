package service

import (
	"MatchingAPI/dto"
	"MatchingAPI/exception"
	"MatchingAPI/service"
	"MatchingAPI/test/util"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFindDriver(t *testing.T) {
	t.Run("Find driver: successfully", func(t *testing.T) {
		matchingService := service.NewMatchingService()
		givenSearchLocationDto := &dto.SearchLocation{
			Radius:    10,
			Latitude:  5.0,
			Longitude: 5.0,
		}

		mockClient := &util.ClientMock{}
		result, err := matchingService.FindDriver(givenSearchLocationDto, mockClient)

		assert.Nil(t, err)
		assert.Equal(t, util.TestMatchResponse.Title, result.Title)
		assert.Equal(t, util.TestMatchResponse.Distance, result.Distance)
	})

	t.Run("Find driver: can not react driver location service", func(t *testing.T) {
		matchingService := service.NewMatchingService()
		givenSearchLocationDto := &dto.SearchLocation{
			Radius:    10,
			Latitude:  5.0,
			Longitude: 5.0,
		}

		mockClient := &util.NilBodyClientMock{}
		result, err := matchingService.FindDriver(givenSearchLocationDto, mockClient)

		assert.Nil(t, result)
		assert.True(t, strings.Contains(err.Error(), "Can not react driver location service"))
	})

	t.Run("Find driver: there is no driver specified location", func(t *testing.T) {
		matchingService := service.NewMatchingService()
		givenSearchLocationDto := &dto.SearchLocation{
			Radius:    10,
			Latitude:  5.0,
			Longitude: 5.0,
		}

		mockClient := &util.ClientDataNotFoundMock{}

		result, err := matchingService.FindDriver(givenSearchLocationDto, mockClient)

		assert.Nil(t, result)
		assert.True(t, strings.Contains(err.Error(), exception.GetDataNotFoundException().Message))
	})

	t.Run("Find driver: Unauthorized request", func(t *testing.T) {
		matchingService := service.NewMatchingService()
		givenSearchLocationDto := &dto.SearchLocation{
			Radius:    10,
			Latitude:  5.0,
			Longitude: 5.0,
		}

		mockClient := &util.ClientUnauthorizedMock{}

		result, err := matchingService.FindDriver(givenSearchLocationDto, mockClient)

		assert.Nil(t, result)
		assert.True(t, strings.Contains(err.Error(), exception.GetUnauthorizedException().Message))
	})
}
