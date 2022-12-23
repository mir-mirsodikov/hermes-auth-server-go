package application

import (
	"context"
	"errors"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/db"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/exception"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/provider"
)

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type CreateUserResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	AccessToken string `json:"accessToken"`
}

func CreateUser(user *CreateUserRequest) (*CreateUserResponse, error) {
	ctx := context.Background()

	if existingUser, err := provider.Queries.GetUserByEmailOrUsername(ctx, db.GetUserByEmailOrUsernameParams{
		Email:    user.Email,
		Username: user.Username,
	}); existingUser != (db.User{}) && err == nil {
		var duplicate string

		if existingUser.Email == user.Email {
			duplicate = "email"
		} else if existingUser.Username == user.Username {
			duplicate = "username"
		}

		return nil, &exception.ApplicationError{
			ErrType: exception.BadRequestError,
			Err:     errors.New(duplicate + " already exists"),
		}
	}

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

	accessToken, err := provider.GenerateToken(createdUser.ID.String())

	if err != nil {
		return nil, &exception.ApplicationError{
			ErrType: exception.AuthorizationError,
			Err:     errors.New("unable to generate access token"),
		}
	}

	return &CreateUserResponse{
		ID:          createdUser.ID.String(),
		Username:    createdUser.Username,
		Email:       createdUser.Email,
		Name:        createdUser.Name,
		AccessToken: accessToken,
	}, nil
}
