package router

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(r *gin.RouterGroup) {
	r.POST("", controllers.CreateUser)
}
