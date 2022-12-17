package application

import "errors"

type GetPingResponse struct {
	Message string `json:"message"`
}

func GetPing() GetPingResponse {
	return GetPingResponse{Message: "pong"}
}

func GetPanic() (GetPingResponse, error) {
	return GetPingResponse{}, errors.New("Panic")
}
