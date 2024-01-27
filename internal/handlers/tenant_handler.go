package handlers

import (
	"kv-store/internal/constants"
	"kv-store/internal/helpers"
	"kv-store/internal/models"
	repository "kv-store/internal/repositories"
	"net/http"

	"github.com/google/uuid"
)

type TenantHandler struct {
	tenantRepo repository.TenantRepository
}

func NewTenantHandler(tenantRepo repository.TenantRepository) *TenantHandler {
	return &TenantHandler{tenantRepo: tenantRepo}
}

func (h *TenantHandler) OnboardTenantHandler(w http.ResponseWriter, r *http.Request) {
	var tenantId string = uuid.New().String()
	var tenantSecret string = uuid.New().String()

	tenant := &models.TenantModel{
		ID:     tenantId,
		Secret: tenantSecret,
	}

	err := h.tenantRepo.Save(tenant)

	if err != nil {
		helpers.SendErrorResponse(w, err.Error(), constants.InternalServerError)
		return
	}

	res := &models.OnboardTenantResponse{
		TenantID:     tenantId,
		TenantSecret: tenantSecret,
	}

	w.Header().Set("Authorization", "Bearer "+helpers.GenerateToken(tenantId, tenantSecret))
	helpers.SendSuccessResponse(w, "Tenant onboarded successfully", res)
	return
}
