package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

// ValidationErrorMessage formats validation errors into a user-friendly format.
func ValidationErrorMessage(err error, v interface{}) map[string]string {
	// Initialize a map to store field-specific error messages
	errors := make(map[string]string)

	// Type assertion to check if the error is of type validator.ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {


		for _, fieldError := range validationErrors {

			// Customize the error messages for each field
			switch fieldError.Tag() {
			case "required":
				errors[fieldError.Field()] = fmt.Sprintf("%s is required", fieldError.Field())
			case "min":
				errors[fieldError.Field()] = fmt.Sprintf("%s must be at least %s characters", fieldError.Field(), fieldError.Param())
			case "max":
				errors[fieldError.Field()] = fmt.Sprintf("%s cannot be longer than %s characters", fieldError.Field(), fieldError.Param())
			case "email":
				errors[fieldError.Field()] = "Invalid email address"
			default:
				errors[fieldError.Field()] = fmt.Sprintf("%s is not valid", fieldError.Field())
			}
		}
	}
	return errors
}
