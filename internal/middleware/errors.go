package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors[0]
			code := 0

			errType := strings.Split(err.Error(), ":")[0]

			fmt.Println(errType)

			switch errType {
			case "Validation Request":
				code = 422
			case "Bad Request":
				code = 400
			case "Authorization":
				code = 401
			default:
				code = 500
			}

			c.JSON(code, gin.H{
				"status":  code,
				"message": err.Error(),
			})
		}
	}
}
