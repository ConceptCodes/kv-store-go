package helpers

import (
	"encoding/json"
	"kv-store/internal/constants"
	"kv-store/internal/models"
	"kv-store/pkg/logger"
	"net/http"
)

func SendSuccessResponse(w http.ResponseWriter, message string, data interface{}) {
	response := models.Response{
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

func SendErrorResponse(w http.ResponseWriter, message string, errorCode string, err error) {

	log := logger.GetLogger()

	log.Error().Err(err).Msg(message)

	response := models.Response{
		Message:   message,
		ErrorCode: errorCode,
	}

	json.NewEncoder(w).Encode(response)
	switch errorCode {
	case constants.NotFound:
		w.WriteHeader(http.StatusNotFound)
	case constants.Unauthorized:
		w.WriteHeader(http.StatusUnauthorized)
	case constants.InternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
	case constants.Forbidden:
		w.WriteHeader(http.StatusForbidden)
	case constants.BadRequest:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
}
