package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomRecoveryMiddleware() func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		fmt.Printf("type of recovered: %T\n", recovered)
		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
