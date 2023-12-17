package middleware

import (
	"github.com/Chandan050/blogwebsite/util"
	"github.com/gin-gonic/gin"
)

func IsAunthenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		if _, err := util.Parsejwt(cookie); err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": "unauthenticated"})
		}
		c.Next()
	}

}
