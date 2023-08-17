package middlewares

import (
	"auth_audit/internal/app/server/services/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware(validateToken interfaces.ValidateToken) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("X-Token")
		if err := validateToken.ValidateToken(token); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Next()
	}
}
