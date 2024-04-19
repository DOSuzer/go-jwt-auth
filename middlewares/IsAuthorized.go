package middlewares

import (
	"os"
	"strings"

	"github.com/DOSuzer/go-jwt-auth/utils"

	"github.com/gin-gonic/gin"
)

var secret = os.Getenv("SECRET_KEY")

func IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")

		if len(t) != 2 {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		token := t[1]

		if token == "" {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		claims, err := utils.ParseToken(token, secret)

		if err != nil {
			c.JSON(401, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		c.Set("role", claims.Role)
		c.Set("email", claims.Subject)
		c.Next()
	}
}
