package exception

type hystrixException struct {
	Message string `json:"message"`
}

func (h *hystrixException) Error() string {
	return h.Message
}

func GetHystrixException() *hystrixException{
	return &hystrixException{
		Message: "HystrixException",
	}
}

type dataNotFoundException struct {
	Message string `json:"message"`
}

func (h *dataNotFoundException) Error() string {
	return h.Message
}

func GetDataNotFoundException() *dataNotFoundException{
	return &dataNotFoundException{
		Message: "Data not found",
	}
}

type genericException struct {
	Message string `json:"message"`
}

func (h *genericException) Error() string {
	return h.Message
}

func GetGenericException(errorMessage string) *genericException{
	return &genericException{
		Message: errorMessage,
	}
}


type unauthorizedException struct {
	Message string `json:"message"`
}

func (h *unauthorizedException) Error() string {
	return h.Message
}

func GetUnauthorizedException() *unauthorizedException{
	return &unauthorizedException{
		Message: "Unauthorized request",
	}
}