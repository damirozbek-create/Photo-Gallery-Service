package middleware

import (
	"net/http"
	"strings"

	"photo-gallery/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr := c.GetHeader("Authorization")

		if tokenStr == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token"})
			c.Abort()
			return
		}

		tokenStr = strings.Replace(tokenStr, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return utils.Secret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		// 🔥 ВАЖНО: безопасный cast
		userID := uint(claims["user_id"].(float64))

		c.Set("user_id", userID)

		c.Next()
	}
}
