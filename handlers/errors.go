package handlers

import (
	"encoding/json"
	"net/http"
)

// APIError represents a custom api error structure
type APIError struct {
	Message string `json:"error"`
}

func writeError(w http.ResponseWriter, errMsg APIError, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(errMsg); err != nil {
		http.Error(w, "Failed to encode error message", http.StatusInternalServerError)
	}
}

func writeSuccess(w http.ResponseWriter, resp interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "Failed to encode resp", http.StatusInternalServerError)
	}
}
