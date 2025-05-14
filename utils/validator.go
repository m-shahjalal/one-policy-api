package utils

import "github.com/go-playground/validator/v10"

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationErrors struct {
	Errors []ValidationError `json:"errors"`
}

var validate *validator.Validate
