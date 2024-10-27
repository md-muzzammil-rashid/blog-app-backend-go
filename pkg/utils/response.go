package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type ApiError struct {
	StatusCode int `json:"status_code"`
	Message string `json:"message"`
	Success bool `json:"success" default:"false"`
	Error []string `json:"error"`
	Data interface{} `json:"data,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface {}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func ReadJSON(r http.Request, data interface{}) error{
	return json.NewDecoder(r.Body).Decode(&data)
}


func WriteError(w http.ResponseWriter, statusCode int, message string) {
	apiError := ApiError{
		StatusCode: statusCode,
		Message: message,
		Success: false,
	}
	WriteJSON(w, statusCode, apiError)
}

func ValidationsError (err validator.ValidationErrors) ApiError{
	apiError := ApiError{
		StatusCode: http.StatusBadRequest,
        Message: "Validation errors",
        Success: false,
	}

	for _, e := range err {
		apiError.Error = append(apiError.Error, e.Field() + ": " + e.ActualTag())
	}
	return apiError
}