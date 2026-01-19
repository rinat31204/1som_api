package middleware

import (
	"1som/internal/application/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ContextUserKey = "auth_user"

func NewAuthMiddleware(provider auth.Provider) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "authorization header required",
			})
			return
		}

		user, err := provider.VerifyToken(
			c.Request.Context(),
			authHeader,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		c.Set(ContextUserKey, user)

		c.Next()
	}
}
