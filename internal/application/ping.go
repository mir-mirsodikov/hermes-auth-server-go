package application

import (
	"errors"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/exception"
)

type GetPingResponse struct {
	Message string `json:"message"`
}

func GetPing() GetPingResponse {
	return GetPingResponse{Message: "pong"}
}

func GetPanic() (GetPingResponse, error) {
	return GetPingResponse{}, &exception.ApplicationError{ErrType: exception.BadRequestError, Err: errors.New("Panic")}
}
