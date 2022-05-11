package service

import (
	"MatchingAPI/dto"
	"MatchingAPI/exception"
	"MatchingAPI/util"
	"encoding/json"
	"github.com/afex/hystrix-go/hystrix"
	"io/ioutil"
	"net/http"
	"sync"
)

var lock = &sync.Mutex{}

var singleInstance IMatchingService

type IMatchingService interface {
	FindDriver(*dto.SearchLocation, util.HttpClient) (*dto.MatchResponse, error)
}

type MatchingService struct{}

func NewMatchingService() IMatchingService {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = &MatchingService{}
		}
	}
	return singleInstance
}

func (m MatchingService) FindDriver(searchLocation *dto.SearchLocation, client util.HttpClient) (*dto.MatchResponse, error) {
	successResponse := make(chan *dto.MatchResponse, 1)
	//errHystrix := hystrix.Go(route.GetRoute(route.SEARCHING), func() error {
	// hystrix.Go asynchronous , hystrix.Do synchronous
	errHystrix := hystrix.Go("/search/driver", func() error {
		request, err := util.FindDriverRequest(searchLocation)
		if err != nil {
			return err
		}

		response, err := client.Do(request)
		if response == nil {
			return exception.GetGenericException("Can not react driver location service")
		}
		defer response.Body.Close()
		if response.StatusCode == http.StatusInternalServerError {
			return exception.GetDataNotFoundException()
		}
		if response.StatusCode == http.StatusUnauthorized {
			return exception.GetUnauthorizedException()
		}
		if err != nil {
			return err
		}

		body, _ := ioutil.ReadAll(response.Body)
		var result = dto.MatchResponse{}

		err = json.Unmarshal(body, &result)
		if err != nil {
			return err
		}

		successResponse <- &result
		return nil
	}, func(err error) error {
		return err
	})

	select {
	case response := <-successResponse:
		return response, nil
	case err := <-errHystrix:
		return nil, err
	}
}