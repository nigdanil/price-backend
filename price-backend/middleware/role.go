package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ManagerOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "manager" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Manager access required"})
			return
		}
		c.Next()
	}
}

func ClientOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "client" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Client access required"})
			return
		}
		c.Next()
	}
}
