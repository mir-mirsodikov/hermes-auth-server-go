package controllers

import (
	"errors"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/application"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/err"
	"github.com/Hermes-chat-App/hermes-auth-server/internal/model"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var requestedUser model.User

	if e := c.ShouldBindJSON(&requestedUser); e != nil {
		c.Errors = append(c.Errors, c.Error(&err.ApplicationError{ErrType: err.BadRequestError, Err: errors.New("invalid request")}))
	}

	user, err := application.CreateUser(&requestedUser)

	if err != nil {
		c.Errors = append(c.Errors, c.Error(err))
	}

	c.JSON(200, user)
}
