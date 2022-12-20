package application

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/gateway"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/model"
)

func CreateUser(user *model.User) (*model.User, error) {
	createdUser, err := gateway.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
