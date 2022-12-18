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

func GetPanic(c *gin.Context) {
	_, err := application.GetPanic()

	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
	}
}
