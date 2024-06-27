package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"phpguru.net/go-api/utils"
)

func Authorization(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}
	c.Set("userId", userId)
	c.Next()
}
