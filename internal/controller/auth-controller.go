package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"github.com/m-shahjalal/onepolicy-api/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

func (ctrl *AuthController) Register(c *gin.Context) {
	type RegisterRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8"`
		Name     string `json:"name" binding:"required"`
	}

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var existingUser model.User
	result := config.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already registered",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process request",
		})
		return
	}

	user := model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Name:     req.Name,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"name":       user.Name,
			"created_at": user.CreatedAt,
		},
	})
}

func (ctrl *AuthController) Login(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user model.User
	result := config.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate tokens",
		})
		return
	}

	// Set cookies
	maxAge := int(time.Hour * 24 * 7 / time.Second) // 7 days
	c.SetCookie("access_token", accessToken, maxAge, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, maxAge, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged in",
		"data": gin.H{
			"user": gin.H{
				"id":    user.ID,
				"email": user.Email,
				"name":  user.Name,
			},
			"tokens": gin.H{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		},
	})
}

func (ctrl *AuthController) GetMe(c *gin.Context) {
	user, exists := c.Get(config.UserContextKey)
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Unauthorized access",
			"error":   "Unauthorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User details fetched successfully",
		"data":    user,
	})
}

func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	oldRefreshToken, err := utils.ExtractRefreshToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid refresh token"})
		return
	}

	claims, err := utils.ValidateToken(oldRefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	var user model.User
	if err := config.DB.Where("id = ?", claims.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID, user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	maxAge := int(time.Hour * 24 * 7 / time.Second) // 7 days
	c.SetCookie("access_token", accessToken, maxAge, "/", "", false, true)
	c.SetCookie("refresh_token", refreshToken, maxAge, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Tokens refreshed successfully",
		"data": gin.H{
			"tokens": gin.H{
				"access_token":  accessToken,
				"refresh_token": refreshToken,
			},
		},
	})
}

func (ctrl *AuthController) Logout(c *gin.Context) {
	c.SetCookie("access_token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
