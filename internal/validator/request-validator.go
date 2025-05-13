package validator

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Use a single instance of Validate, it caches struct info
var validate = validator.New()

// Common validation error response
func ValidationResponse(c *gin.Context, err error) {
	// Cast the error to validator.ValidationErrors
	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request format",
		})
		return
	}

	// Create a map of field errors
	errors := make(map[string]string)
	for _, e := range validationErrors {
		fieldName := e.Field()

		// Customize error message based on validation tag
		switch e.Tag() {
		case "required":
			errors[fieldName] = fieldName + " is required"
		case "email":
			errors[fieldName] = fieldName + " must be a valid email address"
		case "min":
			errors[fieldName] = fieldName + " must be at least " + e.Param() + " characters long"
		case "eqfield":
			errors[fieldName] = fieldName + " must match " + e.Param()
		default:
			errors[fieldName] = e.Error()
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"error":  "Validation failed",
		"errors": errors,
	})
}

// Custom validators
func RegisterCustomValidators() {
	// Strong password validation
	_ = validate.RegisterValidation("strongPassword", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()

		// At least 8 characters, including uppercase, lowercase, number, and special character
		hasUppercase := regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLowercase := regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial := regexp.MustCompile(`[^A-Za-z0-9]`).MatchString(password)

		return len(password) >= 8 && hasUppercase && hasLowercase && hasNumber && hasSpecial
	})
}

// GetValidator returns the validator instance
func GetValidator() *validator.Validate {
	return validate
}
