package helpers

import (
	"encoding/json"
	"kv-store/pkg/models"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := models.Response{
		Message: message,
		Data:    data,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func SendErrorResponse(w http.ResponseWriter, message string, errorCode string) {

	response := models.Response{
		Message:   message,
		ErrorCode: errorCode,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
	switch errorCode {
	case "KV-404":
		w.WriteHeader(http.StatusNotFound)
	case "KV-401":
		w.WriteHeader(http.StatusUnauthorized)
	case "KV-403":
		w.WriteHeader(http.StatusForbidden)
	case "KV-500":
		w.WriteHeader(http.StatusInternalServerError)
	}
}
