package contextkey

// ContextKey is a type for context keys to avoid collisions between packages
type ContextKey string

const (
	// UserIdKey is the context key for user ID
	UserIdKey ContextKey = "user_id"
	// ClaimsKey is the context key for JWT claims
	ClaimsKey ContextKey = "claims"
	// UserKey is the context key for user object
	UserKey ContextKey = "user"
)
