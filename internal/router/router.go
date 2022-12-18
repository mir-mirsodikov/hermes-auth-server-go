package router

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func InitRouter() {
	setupMiddleware()
	getRoutes()
	router.Run()
}

func getRoutes() {
	InitPingRoutes(router.Group("/ping"))
}

func setupMiddleware() {
	router.Use(gin.CustomRecovery(controllers.CustomRecoveryMiddleware()))
	router.Use(controllers.ErrorHandler())
}
