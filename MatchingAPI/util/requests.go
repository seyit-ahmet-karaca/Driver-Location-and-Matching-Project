package util

import (
	"MatchingAPI/config"
	"MatchingAPI/dto"
	"fmt"
	"net/http"
)

func FindDriverRequest(searchLocation *dto.SearchLocation) (*http.Request, error) {
	url := fmt.Sprintf("http://driver-location:3000/api/search/location?radius=%d&longitude=%.20f&latitude=%.20f",
		searchLocation.Radius,
		searchLocation.Longitude,
		searchLocation.Latitude)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(config.Get().MatchApiHeader, config.Get().SecretMatchApiKey)
	return request, nil
}

