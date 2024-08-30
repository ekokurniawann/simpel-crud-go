package response

import (
	"encoding/json"
	"net/http"
)

func ErrorResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(map[string]interface{}{"error": message}); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}

func SuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "error encoding response", http.StatusInternalServerError)
	}
}
