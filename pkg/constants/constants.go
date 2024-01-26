package constants

const (
	// Error codes
	NotFound            = "ZM-404"
	BadRequest          = "ZM-400"
	Unauthorized        = "ZM-401"
	Forbidden           = "ZM-403"
	InternalServerError = "ZM-500"

	// Queries
	FindByIdQuery = "id = ?"
	FindByTenantIdAndKey = "tenant_id = ? AND id = ?"
)
