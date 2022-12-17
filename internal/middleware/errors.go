package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		fmt.Println("After the request")
		if len(c.Errors) > 0 {
			c.JSON(500, gin.H{
				"status":  500,
				"message": "Internal Server Error",
			})
		}
	}
}