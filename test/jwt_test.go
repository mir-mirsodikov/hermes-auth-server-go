package test

import (
	"strings"
	"testing"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/provider"
)

func TestCreateToken(t *testing.T) {
	userId := "123"
	token, _ := provider.GenerateToken(userId)

	if token == "" {
		t.Errorf("GenerateToken() = %q, want %q", token, "token")
	}

	tokenParts := strings.Split(token, ".")

	if len(tokenParts) != 3 {
		t.Errorf("GenerateToken() = %q, want %q", tokenParts, "tokenParts")
	}
}

func TestVerifyToken(t *testing.T) {
	userId := "123"

	token, _ := provider.GenerateToken(userId)

	verifiedUserId, err := provider.VerifyToken(token)

	if err != nil {
		t.Errorf("VerifyToken() = %q, error", err)
	}

	if verifiedUserId != userId {
		t.Errorf("VerifyToken() = %q, want %q", verifiedUserId, userId)
	}
}
