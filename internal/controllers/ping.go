package controllers

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/application"
	"github.com/gin-gonic/gin"
)

func GetPing(c *gin.Context) {
	response := application.GetPing()

	c.JSON(200, gin.H{
		"message": response.Message,
	})
}