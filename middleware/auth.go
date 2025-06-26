package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sajid414/Go-react-auth/helper"
)

func Authenticate(c *gin.Context) {
	token := c.GetHeader("token")
	if token == "" {
		c.JSON(401, gin.H{"Error": "Token not present."})
		c.Abort()
		return
	}
	claims, msg := helper.ValidateToken(token)
	log.Println(claims)

	if msg != "" {
		c.JSON(401, gin.H{"Error": msg})
		c.Abort()
		return
	}
	c.Next()
}
