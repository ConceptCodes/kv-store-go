package handlers

import (
	"fmt"
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

// OnboardTenantHandler godoc
// @Summary Onboard Tenant
// @Description Onboard Tenant
// @Tags Tenant
// @Accept  json
// @Produce  json
// @Success 200 {object} OnboardTenantResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /tenants/onboard [post]
func (h *TenantHandler) OnboardTenantHandler(w http.ResponseWriter, r *http.Request) {
	tenantId := uuid.New().String()
	tenantSecret := uuid.New().String()

	tenant := &models.TenantModel{
		ID:     tenantId,
		Secret: tenantSecret,
	}

	err := h.tenantRepo.Save(tenant)

	if err != nil {
		message := fmt.Sprintf(constants.SaveEntityError, "Tenant")
		helpers.SendErrorResponse(w, message, constants.InternalServerError, err)
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
