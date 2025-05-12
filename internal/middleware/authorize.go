package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthorizeRole(requiredRole int8) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Нет доступа"})
			return
		}

		role, ok := roleVal.(int8)
		if !ok || role < requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав"})
			return
		}

		c.Next()
	}
}
