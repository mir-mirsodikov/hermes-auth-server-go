package application

import (
	"context"
	"errors"
	"fmt"
	"log"

	"time"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/db"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/exception"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/provider"
	"github.com/google/uuid"
)

type VerifyCodeRequest struct {
	ID   uuid.UUID `json:"id" binding:"required"`
	Code int       `json:"code" binding:"required"`
}

type VerifyCodeResponse struct {
	Valid bool `json:"valid"`
}

func VerifyCode(r *VerifyCodeRequest) (*VerifyCodeResponse, error) {
	ctx := context.Background()

	userCode, err := provider.Queries.GetVerificationByUser(ctx, r.ID)

	if err != nil {
		return &VerifyCodeResponse{Valid: false}, &exception.ApplicationError{
			ErrType: exception.BadRequestError,
			Err:     errors.New("unable to get verification code"),
		}
	}

	if time.Now().UTC().After(userCode.CreatedAt.Add(5 * time.Minute)) {
		return &VerifyCodeResponse{Valid: false}, &exception.ApplicationError{
			ErrType: exception.BadRequestError,
			Err:     errors.New("verification code expired"),
		}
	}

	if r.Code != int(userCode.Code) {
		return &VerifyCodeResponse{Valid: false}, nil
	}

	return &VerifyCodeResponse{Valid: true}, nil
}

type LoginRequest struct {
	Email string `json:"email" binding:"required"`
}

type LoginResponse struct {
	ID uuid.UUID `json:"id"`
}

func Login(r *LoginRequest) (*LoginResponse, error) {
	ctx := context.Background()

	foundUser, err := provider.Queries.GetUserByEmail(ctx, r.Email)

	if err != nil {
		return nil, &exception.ApplicationError{
			ErrType: exception.BadRequestError,
			Err:     errors.New("user not found"),
		}
	}

	go func() {
		code := provider.GenerateVerificationCode()
		msg := fmt.Sprintf(getLoginEmailMessage(), foundUser.Name, code)
		provider.Queries.CreateVerification(ctx, db.CreateVerificationParams{
			UserID: foundUser.ID,
			Code:   int32(code),
		})

		if err := provider.SendEmail(foundUser.Email, "Subject: Hermes Login Code", msg); err != nil {
			log.Println(err)
		}
	}()

	return &LoginResponse{
		ID: foundUser.ID,
	}, nil
}

func getLoginEmailMessage() string {
	return `
Welcome back to Hermes, %s!

Please enter the following code to login to your account: %d
This code will expire in 5 minutes.

If you did not request this code, please ignore this email.

Please do not reply to this email.
	`
}
