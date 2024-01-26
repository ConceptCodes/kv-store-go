package constants

const (
	// Error codes
	NotFound            = "KV-404"
	BadRequest          = "KV-400"
	Unauthorized        = "KV-401"
	Forbidden           = "KV-403"
	InternalServerError = "KV-500"

	// Messages
	EntityNotFound = "%s with id %s not found"

	// Queries
	FindByIdQuery             = "id = ?"
	FindByTenantIdAndKeyQuery = "tenant_id = ? AND id = ?"
	FindExpiredRecordsQuery   = "expires_at < NOW()"
)
