package util

import (
	"MatchingAPI/dto"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var TestMatchResponse = &dto.MatchResponse{
	Distance: 10,
	Title:    "Test",
}

type ClientMock struct {
}

func (c *ClientMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		Body: io.NopCloser(strings.NewReader(fmt.Sprintf("{\"distance\":%f, \"title\":\"%s\"}",
			TestMatchResponse.Distance, TestMatchResponse.Title))),
	}, nil
}

type NilBodyClientMock struct {
}

func (c *NilBodyClientMock) Do(req *http.Request) (*http.Response, error) {
	return nil, nil
}

type ClientDataNotFoundMock struct {
}

func (c *ClientDataNotFoundMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusInternalServerError,
		Body: io.NopCloser(strings.NewReader(fmt.Sprintf("{\"distance\":%f, \"title\":\"%s\"}",
			TestMatchResponse.Distance, TestMatchResponse.Title))),
	}, nil
}

type ClientUnauthorizedMock struct {
}

func (c *ClientUnauthorizedMock) Do(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusUnauthorized,
		Body: io.NopCloser(strings.NewReader(fmt.Sprintf("{\"distance\":%f, \"title\":\"%s\"}",
			TestMatchResponse.Distance, TestMatchResponse.Title))),
	}, nil
}

