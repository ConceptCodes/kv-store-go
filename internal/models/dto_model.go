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
	Key     string `json:"key"`
	Value   string `json:"value"`
	Expires string `json:"expires"`
}

type SaveRecordRequest struct {
	Key   string `json:"key" validate:"required,noSQLKeywords"`
	Value string `json:"value" validate:"required,noSQLKeywords"`
	TTL   int    `json:"ttl"`
}
