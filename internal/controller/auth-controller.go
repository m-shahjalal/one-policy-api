package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/m-shahjalal/onepolicy-api/config"
	"github.com/m-shahjalal/onepolicy-api/internal/model"
	"golang.org/x/crypto/bcrypt"
)

// AuthController handles authentication related requests
type AuthController struct{}

// Register handles user registration
func (ctrl *AuthController) Register(c *gin.Context) {
	// Example request structure
	type RegisterRequest struct {
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required,min=8"`
		First_name string `json:"first_name" binding:"required"`
		Last_name  string `json:"last_name" binding:"required"`
	}

	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Check if email already exists
	var existingUser model.User
	result := config.DB.Where("email = ?", req.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusConflict, gin.H{
			"error": "Email already registered",
		})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process request",
		})
		return
	}

	// Create new user
	user := model.User{
		Email:      req.Email,
		Password:   string(hashedPassword),
		First_name: req.First_name,
		Last_name:  req.Last_name,
	}

	// Save user to database
	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to register user",
		})
		return
	}

	// Return response without password
	user.Password = ""
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"first_name": user.First_name,
			"last_name":  user.Last_name,
			"created_at": user.CreatedAt,
		},
	})
}

// Login authenticates a user and returns a token
func (ctrl *AuthController) Login(c *gin.Context) {
	type LoginRequest struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		println("error", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the user by email
	var user model.User
	result := config.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare the provided password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Generate tokens (in a real app, you'd use a JWT library)
	accessToken := "access-token-" + user.Email + "-" + time.Now().String()   // Replace with proper JWT token
	refreshToken := "refresh-token-" + user.Email + "-" + time.Now().String() // Replace with proper JWT token

	// Store the refresh token in the database
	user.Token = accessToken
	user.Refresh_token = refreshToken

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process login",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"first_name": user.First_name,
			"last_name":  user.Last_name,
		},
	})
}

// Logout invalidates user tokens
func (ctrl *AuthController) Logout(c *gin.Context) {
	// Extract token from Authorization header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No authorization token provided",
		})
		return
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Find user with this token
	var user model.User
	result := config.DB.Where("token = ?", token).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	// Clear tokens
	user.Token = ""
	user.Refresh_token = ""

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process logout",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Logged out successfully",
	})
}

// RefreshToken issues a new access token using a refresh token
func (ctrl *AuthController) RefreshToken(c *gin.Context) {
	// Example request structure
	type RefreshRequest struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	var req RefreshRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find user with this refresh token
	var user model.User
	result := config.DB.Where("refresh_token = ?", req.RefreshToken).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid refresh token",
		})
		return
	}

	// Generate new tokens (in a real app, you'd use a JWT library)
	newAccessToken := "access-token-" + user.Email + "-" + time.Now().String()   // Replace with proper JWT token
	newRefreshToken := "refresh-token-" + user.Email + "-" + time.Now().String() // Replace with proper JWT token

	// Update tokens in database
	user.Token = newAccessToken
	user.Refresh_token = newRefreshToken

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to refresh token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Token refreshed successfully",
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}

// ForgotPassword initiates the password reset process
func (ctrl *AuthController) ForgotPassword(c *gin.Context) {
	// Example request structure
	type ForgotPasswordRequest struct {
		Email string `json:"email" binding:"required,email"`
	}

	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find the user by email
	var user model.User
	result := config.DB.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		// Don't reveal that the email doesn't exist for security reasons
		c.JSON(http.StatusOK, gin.H{
			"message": "If your email is registered, you will receive a password reset link",
		})
		return
	}

	// Generate a reset token (in a real app, this would be a cryptographically secure token)
	resetToken := "reset-token-" + user.Email + "-" + time.Now().String()

	// In a real application, you would:
	// 1. Store this token in the database (possibly in a separate password_resets table)
	// 2. Set an expiration time for the token
	// 3. Send an email with a link containing this token

	// For this example, we'll just update the user's token field
	user.Token = resetToken
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process request",
		})
		return
	}

	// In a real app, you would send an email here
	// sendResetEmail(user.Email, resetToken)

	c.JSON(http.StatusOK, gin.H{
		"message": "If your email is registered, you will receive a password reset link",
		// For development/testing purposes only - remove in production
		"reset_token": resetToken,
	})
}

// ResetPassword completes the password reset process
func (ctrl *AuthController) ResetPassword(c *gin.Context) {
	// Example request structure
	type ResetPasswordRequest struct {
		Token           string `json:"token" binding:"required"`
		Password        string `json:"password" binding:"required,min=8"`
		ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
	}

	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Find user with this reset token
	var user model.User
	result := config.DB.Where("token = ?", req.Token).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid or expired reset token",
		})
		return
	}

	// In a real application, you would check if the token is expired
	// if isTokenExpired(req.Token) { ... }

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to process request",
		})
		return
	}

	// Update password and clear the reset token
	user.Password = string(hashedPassword)
	user.Token = "" // Clear the reset token

	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to reset password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password has been reset successfully",
	})
}

// GetMe retrieves the current authenticated user's information
func (ctrl *AuthController) GetMe(c *gin.Context) {
	// Extract token from Authorization header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No authorization token provided",
		})
		return
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Find user with this token
	var user model.User
	result := config.DB.Where("token = ?", token).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	// Return user data without sensitive information
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"email":      user.Email,
		"first_name": user.First_name,
		"last_name":  user.Last_name,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	})
}

// UpdateProfile updates the current user's profile information
func (ctrl *AuthController) UpdateProfile(c *gin.Context) {
	// Example request structure
	type UpdateProfileRequest struct {
		First_name string `json:"first_name"`
		Last_name  string `json:"last_name"`
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Extract token from Authorization header
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No authorization token provided",
		})
		return
	}

	// Remove "Bearer " prefix if present
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// Find user with this token
	var user model.User
	result := config.DB.Where("token = ?", token).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	// Update user fields if provided
	if req.First_name != "" {
		user.First_name = req.First_name
	}

	if req.Last_name != "" {
		user.Last_name = req.Last_name
	}

	// Save changes to database
	if err := config.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update profile",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile updated successfully",
		"user": gin.H{
			"id":         user.ID,
			"email":      user.Email,
			"first_name": user.First_name,
			"last_name":  user.Last_name,
			"updated_at": user.UpdatedAt,
		},
	})
}
