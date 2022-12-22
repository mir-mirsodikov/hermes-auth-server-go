package application

import (
	"context"
	"errors"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/db"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/exception"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/model"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/provider"
)

func CreateUser(user *model.User) (*model.User, error) {
	ctx := context.Background()
	createdUser, err := provider.Queries.CreateUser(ctx, db.CreateUserParams{
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
	})

	if err != nil {
		return nil, &exception.ApplicationError{
			ErrType: exception.BadRequestError,
			Err:     errors.New("unable to create user in the database"),
		}
	}

	return &model.User{
		Id:       createdUser.ID.String(),
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Name:     createdUser.Name,
	}, nil
}
