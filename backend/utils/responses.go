package utils

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, statusCode int, message string, errorFlag bool) {
	resp := map[string]interface{}{
		"error":   errorFlag,
		"message": message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}


func SendErrorResponse(w http.ResponseWriter, statusCode int,message string) {
	if statusCode < 400 || statusCode > 500 {
		statusCode = http.StatusInternalServerError
	}
	SendResponse(w, statusCode, message, true)
}

func SendSuccessResponse(w http.ResponseWriter, statusCode int, message string) {
	if statusCode < 200 || statusCode > 299 {
		statusCode = http.StatusOK
	}
	SendResponse(w, statusCode, message, false)
}