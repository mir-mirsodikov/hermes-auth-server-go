package main

import (
	"log"
	"os"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/provider"
	route "github.com/Hermes-chat-App/hermes-auth-server/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	connStr := os.Getenv("DATABASE_URL")
	tokenSecret := os.Getenv("TOKEN_SECRET")

	if port == "" {
		port = "8080"
	}

	if connStr == "" {
		connStr = "postgresql://postgres@localhost:5432/user_db"
	}

	if tokenSecret == "" {
		tokenSecret = "secret"
	}

	provider.Init(connStr, tokenSecret)
	route.InitRouter()
}
