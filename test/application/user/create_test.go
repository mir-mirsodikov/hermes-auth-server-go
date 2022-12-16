package test

import (
	"testing"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/application/user"
)

func TestCreateUser(t *testing.T) {
	name := "Jane Doe"
	want := "Created user: " + name + "!"

	if got := user.CreateUser(name); got != want {
		t.Errorf("CreateUser(%q) = %q, want %q", name, got, want)
	}
}
