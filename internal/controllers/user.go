package controllers

import (
	"github.com/Hermes-chat-App/hermes-auth-server/internal/application"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var requestedUser application.CreateUserRequest

	if !ValidateRequest(c, &requestedUser) {
		return
	}

	user, err := application.CreateUser(&requestedUser)

	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
		return
	}

	c.JSON(200, user)
}
