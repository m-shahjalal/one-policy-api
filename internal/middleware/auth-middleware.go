package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"github.com/m-shahjalal/onepolicy-api/utils"
)

func getUserFromClaims(claims *utils.TokenClaims) (any, error) {
	var user model.User
	if err := config.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		return nil, err
	}

	userData := map[string]any{
		"id":         user.ID,
		"email":      user.Email,
		"name":       user.Name,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	return userData, nil
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.ExtractAccessToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(token)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set(config.IsAuthContextKey, true)
		c.Set(config.TokenContextKey, token)
		c.Set(config.ClaimsContextKey, claims)
		c.Set(config.UserIDContextKey, claims.UserID)

		user, err := getUserFromClaims(claims)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}
		c.Set(config.UserContextKey, user)

		c.Next()
	}
}
