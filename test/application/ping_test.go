package test

import (
	"testing"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/application"
)

func TestGetPing(t *testing.T) {
	want := application.GetPingResponse{
		Message: "pong",
	}

	if got := application.GetPing(); got != want {
		t.Errorf("GetPing() = %q, want %q", got, want)
	}
}
