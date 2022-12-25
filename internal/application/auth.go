package application

import (
	"context"
	"errors"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/exception"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/provider"
	"github.com/google/uuid"
	"time"
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

	if time.Now().After(userCode.CreatedAt.Time.Add(5 * time.Minute)) {
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
