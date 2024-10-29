package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Success bool `json:"success" default:"false"`
	Error []string `json:"error"`
	Data interface{} `json:"data,omitempty"`
}

type ApiResponse struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Success bool `json:"success" default:"true"`
	Data interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, message string, data interface {}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(ApiResponse{StatusCode: statusCode, Message: message, Data: data})
}

func ReadJSON(r http.Request, data interface{}) error{
	return json.NewDecoder(r.Body).Decode(&data)
}


func WriteError(w http.ResponseWriter, statusCode int, message string) {
	// apiError := ApiError{
	// 	StatusCode: statusCode,
	// 	Message: message,
	// 	Success: false,
	// }
	WriteJSON(w, statusCode, message, nil)
}

// ValidationError implements error interface
type ValidationError struct {
	Errors map[string]string
}

func (v ValidationError) Error() string {
	var errMsgs []string
	for field, message := range v.Errors {
		errMsgs = append(errMsgs, fmt.Sprintf("%s: %s", field, message))
	}
	return strings.Join(errMsgs, ", ")
}

func ValidationsError (errs validator.ValidationErrors) error{
	validationErrors := make(map[string]string)

	for _, err := range errs {
		// Build a user-friendly error message
		var errorMessage string
		switch err.Tag() {
		case "required":
			errorMessage = fmt.Sprintf("%s is required", err.Field())
		case "email":
			errorMessage = fmt.Sprintf("%s must be a valid email address", err.Field())
		case "min":
			errorMessage = fmt.Sprintf("%s must be at least %s characters", err.Field(), err.Param())
		case "max":
			errorMessage = fmt.Sprintf("%s cannot be more than %s characters", err.Field(), err.Param())
		default:
			errorMessage = fmt.Sprintf("%s is not valid", err.Field())
		}

		// Add the error message to the map with the field name as the key
		validationErrors[err.Field()] = errorMessage
	}

	return ValidationError{Errors: validationErrors}
}