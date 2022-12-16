package main

import (
	"fmt"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/application"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/application/user"
)

func main() {
  fmt.Println("Hello, world")
	fmt.Println(application.SayHello("John Doe"))
	fmt.Println(user.CreateUser("Jane Doe"))
}
