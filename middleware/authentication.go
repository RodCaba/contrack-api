package middleware

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get token from header
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
		c.Abort()
	}

	// Verify token
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if token is expired
		if claims["exp"].(float64) < float64(time.Now().Unix()) {
			c.JSON(401, gin.H{"error": "Token expired!"})
			c.Abort()
		}
		c.Set("sub", claims["sub"])
		c.Next()
	} else {
		c.JSON(401, gin.H{"error": "Unauthorized!"})
		c.Abort()
	}
}