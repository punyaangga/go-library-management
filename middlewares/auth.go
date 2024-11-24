package middlewares

import (
	"libraryManagement/config"
	"libraryManagement/models"
	"libraryManagement/utils"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware validates the JWT token
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Authorization header is required", nil)
			c.Abort()
			return
		}

		// Parse token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Check if token is expired
		var expiredToken models.ExpiredToken
		if err := config.DB.Where("token = ?", tokenString).First(&expiredToken).Error; err == nil {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Token has been revoked", nil)
			c.Abort()
			return
		}

		// Parse and validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			utils.SendErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token", nil)
			c.Abort()
			return
		}

		// Pass the user information to the next handler
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			c.Set("username", claims["username"])
		}

		c.Next()
	}
}
