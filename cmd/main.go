package main

import (
	"log"
	"os"

	route "github.com/Hermes-chat-App/hermes-auth-server/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	route.InitRouter()
}
