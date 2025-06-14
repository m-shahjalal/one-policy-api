package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
)

// GetCurrentUser retrieves the user from context
func GetCurrentUser(c *gin.Context) (interface{}, bool) {
	user, exists := c.Get(config.UserContextKey)
	return user, exists
}

// GetCurrentUserID retrieves the user ID from context
func GetCurrentUserID(c *gin.Context) (uint, bool) {
	userID, exists := c.Get(config.UserIDContextKey)
	if !exists {
		return 0, false
	}

	if id, ok := userID.(uint); ok {
		return id, true
	}
	return 0, false
}

// GetClaims retrieves the JWT claims from context
func GetClaims(c *gin.Context) (*TokenClaims, bool) {
	claims, exists := c.Get(config.ClaimsContextKey)
	if !exists {
		return nil, false
	}

	if authClaims, ok := claims.(*TokenClaims); ok {
		return authClaims, true
	}
	return nil, false
}

func GetToken(c *gin.Context) (string, bool) {
	token, exists := c.Get(config.TokenContextKey)
	if !exists {
		return "", false
	}

	if tokenStr, ok := token.(string); ok {
		return tokenStr, true
	}
	return "", false
}

func IsAuthenticated(c *gin.Context) bool {
	isAuth, exists := c.Get(config.IsAuthContextKey)
	if !exists {
		return false
	}

	if auth, ok := isAuth.(bool); ok {
		return auth
	}
	return false
}

func MustGetCurrentUser(c *gin.Context) interface{} {
	user, exists := GetCurrentUser(c)
	if !exists {
		panic("User not found in context")
	}
	return user
}

// MustGetCurrentUserID retrieves user ID or panics (use carefully)
func MustGetCurrentUserID(c *gin.Context) uint {
	userID, exists := GetCurrentUserID(c)
	if !exists {
		panic("User ID not found in context")
	}
	return userID
}
