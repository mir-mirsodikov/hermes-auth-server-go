package middleware

import "github.com/gin-gonic/gin"

func CustomRecoveryMiddleware() func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
	}
}
