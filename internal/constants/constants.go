package constants

const (
	// Error codes
	NotFound            = "KV-404"
	BadRequest          = "KV-400"
	Unauthorized        = "KV-401"
	Forbidden           = "KV-403"
	InternalServerError = "KV-500"

	// Messages
	EntityNotFound  = "%s with id %s does not exist."
	SaveEntityError = "Error while saving %s."

	// Queries
	FindByIdQuery             = "id = ?"
	FindByTenantIdAndKeyQuery = "tenant_id = ? AND id = ?"
	FindExpiredRecordsQuery   = "expires_at < DATETIME('now')"

	// Misc
	CronDelayInSeconds = 5
	TimeFormat         = "2006-01-02 15:04:05"
	TraceIdHeader      = "X-Trace-Id"
)
