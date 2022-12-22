package controllers

import (
	"errors"
	"log"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/exception"
	"github.com/gin-gonic/gin"
)

// ValidateRequest validates the request body
// and returns true if the request is valid
func ValidateRequest[R any](c *gin.Context, r *R) bool {
	if e := c.ShouldBindJSON(r); e != nil {
		log.Println(e)
		c.Errors = append(c.Errors, c.Error(&exception.ApplicationError{ErrType: exception.BadRequestError, Err: errors.New("invalid request")}))
		return false
	}

	return true
}
