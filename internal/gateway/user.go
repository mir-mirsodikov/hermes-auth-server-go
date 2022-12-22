package gateway

import (
	"context"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/db"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/model"
)

func CreateUser(user *model.User) (*model.User, error) {
	ctx := context.Background()
	createdUser, err := Queries.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	})

	if err != nil {
		return nil, err
	}

	return &model.User{
		Id:       createdUser.ID.String(),
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Name:     createdUser.Name,
	}, nil
}
