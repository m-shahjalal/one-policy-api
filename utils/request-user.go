package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/m-shahjalal/onepolicy-api/internal/model"
)

// GetUserID retrieves the user ID from the context.
func GetUserID(c *gin.Context) uuid.UUID {
	user, exists := c.Get("user")
	if !exists {
		return uuid.UUID{} // or handle the case where user is not found
	}

	u, ok := user.(*model.User)
	if !ok {
		return uuid.UUID{} // or handle the case where user is not of type *model.User
	}

	return u.ID
}
