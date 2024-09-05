package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

type AuthUser struct {
	ID uint `json:"id" binding:"required"`
}

func IsAuthenticated(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "" {
		token = token[7:]
		token, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			secret := os.Getenv("JWT_SECRET")
			return []byte(secret), nil
		})
		if err == nil {
			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				id := uint(claims["id"].(float64))
				if float64(time.Now().Unix()) < claims["exp"].(float64) {
					c.Set("user", AuthUser{
						ID: id,
					})
					c.Next()
					return
				}
			}
		}
	}
	c.JSON(http.StatusUnauthorized, gin.H{
		"error": true,
		"msg":   "Unauthorized",
	})
	c.Abort()
}
