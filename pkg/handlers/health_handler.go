package handlers

import (
	"net/http"
	"kv-store/pkg/helpers"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServiceAliveHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendSuccessResponse(w, "Service is alive", nil)
}
