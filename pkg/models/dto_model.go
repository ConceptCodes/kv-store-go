package models

type OnboardTenantResponse struct {
	TenantID     string `json:"tenant_id"`
	TenantSecret string `json:"tenant_secret"`
}

type GetRecordRequest struct {
	TenantID string `json:"tenant_id"`
	Key      string `json:"key"`
}

type GetRecordResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SaveRecordRequest struct {
	Key   string `json:"key" validate:"required, regexp=^[a-zA-Z0-9]*$"`
	Value string `json:"value" validate:"required, regexp=^[a-zA-Z0-9]*$"`
	TTL   uint   `json:"ttl"`
}
