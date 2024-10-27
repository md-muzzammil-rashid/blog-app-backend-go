package utils

import (
	"encoding/json"
	"net/http"
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


func WriteError(w http.ResponseWriter, statusCode int, message string) {
	apiError := ApiError{
		StatusCode: statusCode,
		Message: message,
		Success: false,
	}
	WriteJSON(w, statusCode, apiError)
}