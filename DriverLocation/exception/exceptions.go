package exception

import "fmt"

const MongoDbInsertException = "mongoDbInsertException"
const MongoDbReadException = "mongoDbReadException"
const GenericException = "genericException"
const ThereIsNoDriverException = "thereIsNoDriverException"

type mongoDbInsertException struct {
	Message string `json:"message"`
	Type    string `json:"-"`
}

func (m *mongoDbInsertException) Error() string {
	return MongoDbInsertException + " " + m.Message
}

func GetMongoDbInsertException(errorMessage string) *mongoDbInsertException {
	return &mongoDbInsertException{
		Message: errorMessage,
	}
}

type mongoDbReadException struct {
	Message string `json:"message"`
}

func (m *mongoDbReadException) Error() string {
	return MongoDbReadException + " " + m.Message
}

func GetMongoDbReadException(errorMessage string) *mongoDbReadException {
	return &mongoDbReadException{
		Message: errorMessage,
	}
}

type genericException struct {
	Message string `json:"message"`
}

func (m *genericException) Error() string {
	return GenericException + " " + m.Message
}

func GetGenericException(errorMessage string) *genericException {
	return &genericException{
		Message: errorMessage,
	}
}

type thereIsNoDriverException struct {
	Message string `json:"message"`
}

func (m *thereIsNoDriverException) Error() string {
	return ThereIsNoDriverException + " " + m.Message
}

func GetThereIsNoDriverException(radius int) *thereIsNoDriverException {
	return &thereIsNoDriverException{
		Message: fmt.Sprintf("There is no driver in radius :%d km", radius),
	}
}
