package controller

import (
	"MatchingAPI/controller"
	"MatchingAPI/dto"
	"MatchingAPI/mock"
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMatchingController(t *testing.T) {
	t.Run("Search Driver Location", func(t *testing.T) {
		givenLatitude := 0.1
		givenLongitude := 0.1
		givenRadius := 5

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf(
			"/search/driver?longitude=%.20f&latitude=%.20f&radius=%d", givenLongitude, givenLatitude, givenRadius),
			http.NoBody)

		rec := httptest.NewRecorder()
		context := e.NewContext(req, rec)

		mockController := gomock.NewController(t)
		defer mockController.Finish()

		location := &dto.SearchLocation{
			Radius:    givenRadius,
			Latitude:  givenLatitude,
			Longitude: givenLongitude,
		}
		client := &http.Client{}

		givenMatchResponse := &dto.MatchResponse{
			Distance: 10,
			Title:    "Test",
		}

		mockMatchingService := mock.NewMockIMatchingService(mockController)
		mockMatchingService.EXPECT().
			FindDriver(location, client).
			Return(givenMatchResponse, nil).
			Times(1)

		matchingController := controller.NewMatchingController(mockMatchingService)
		result := matchingController.FindDriver(context)

		actual := &dto.MatchResponse{}
		err := json.Unmarshal(rec.Body.Bytes(), actual)

		assert.NoError(t, err)
		assert.NoError(t, result)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, givenMatchResponse.Title, actual.Title)
		assert.Equal(t, givenMatchResponse.Distance, actual.Distance)
	})

}
