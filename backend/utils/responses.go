package utils

import (
	"encoding/json"
	"net/http"
)

func SendErrorResponse(w http.ResponseWriter, message string) {
	resp := map[string]interface{}{
		"error":   true,
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(resp)
}

func SendSuccessResponse(w http.ResponseWriter, message string) {
	resp := map[string]interface{}{
		"error":   false,
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}