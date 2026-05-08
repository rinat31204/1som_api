package middleware

import (
	"1som/internal/auth/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

const ContextUserKey = "uid"

func NewAuthMiddleware(authService application.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		//token := c.GetHeader("Authorization")
		//if token == "" {
		//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		//		"error": "authorization header required",
		//	})
		//	return
		//}
		token := ""
		user, err := authService.VerifyUser(
			c.Request.Context(),
			token,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		c.Set(ContextUserKey, user.UID)

		c.Next()
	}
}
