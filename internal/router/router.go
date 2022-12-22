package router

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func InitRouter() {
	router.SetTrustedProxies(nil)
	setupMiddleware()
	getRoutes()
	router.Run()
}

func getRoutes() {
	InitPingRoutes(router.Group("/ping"))
	InitUserRoutes(router.Group("/user"))
}

func setupMiddleware() {
	router.Use(gin.CustomRecovery(controllers.CustomRecoveryMiddleware()))
	router.Use(controllers.ErrorHandler())
}
