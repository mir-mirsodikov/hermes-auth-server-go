package router

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitPingRoutes(r *gin.RouterGroup) {
	r.GET("", controllers.GetPing)
	r.GET("/panic", controllers.GetPanic)
}
