package router

import (
	"log"
	"os"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router = gin.Default()

func InitRouter() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	setupMiddleware()
	getRoutes()
	router.Run()
}

func getRoutes() {
	InitPingRoutes(router.Group("/ping"))
}

func setupMiddleware() {
	router.Use(gin.CustomRecovery(middleware.CustomRecoveryMiddleware()))
}
