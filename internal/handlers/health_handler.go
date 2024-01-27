package handlers

import (
	"kv-store/internal/helpers"
	"net/http"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// ServiceAliveHandler godoc
// @Summary Check if service is alive
// @Description Check if service is alive
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Router /api/health/alive [get]
func (h *HealthHandler) ServiceAliveHandler(w http.ResponseWriter, r *http.Request) {
	helpers.SendSuccessResponse(w, "Service is alive", nil)
	return
}
