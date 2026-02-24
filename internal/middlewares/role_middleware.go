package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func OnlyADM() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != "ADM" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "Forbidden",
			})
			return
		}
		c.Next()
	}
}
