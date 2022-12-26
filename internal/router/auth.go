package router

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitAuthRoutes(r *gin.RouterGroup) {
	r.POST("/verify", controllers.VerifyCode)
}
