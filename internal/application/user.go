package application

import (
	"context"
	"errors"
	"fmt"
	"log"

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

	go func() {
		code := provider.GenerateVerificationCode()
		msg := fmt.Sprintf(getMessage(), createdUser.Name, code)
		provider.Queries.CreateVerification(ctx, db.CreateVerificationParams{
			UserID: createdUser.ID,
			Code:   int32(code),
		})
		if err := provider.SendEmail(createdUser.Email, "Subject: Welcome to Hermes\n", msg); err != nil {
			log.Println(err)
		}
	}()

	return &CreateUserResponse{
		ID:          createdUser.ID.String(),
		Username:    createdUser.Username,
		Email:       createdUser.Email,
		Name:        createdUser.Name,
		AccessToken: accessToken,
	}, nil
}

func getMessage() string {
	return `
Welcome to Hermes, %s!

Your verification code is %d.
This code will expire in 5 minutes.

Please do not reply to this email.
	`
}
