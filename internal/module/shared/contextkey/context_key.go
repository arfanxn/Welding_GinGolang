package contextkey

import (
	"context"

	"github.com/arfanxn/welding/internal/module/shared/domain/entity"
)

// ContextKey is a type for context keys to avoid collisions between packages
type ContextKey string

const (
	// UserIdKey is the context key for user ID
	UserIdKey ContextKey = "user_id"
	// ClaimsKey is the context key for JWT claims
	ClaimsKey ContextKey = "claims"
	// UserKey is the context key for user object
	UserKey ContextKey = "user"
	// UserPermissionsKey is the context key for user permissions
	UserPermissionsKey ContextKey = "user_permissions"
	// UserAgentKey is the context key for user agent
	UserAgentKey ContextKey = "user_agent"
	// ClientIpKey is the context key for client IP
	ClientIpKey ContextKey = "client_ip"
	// RequestURLKey is the context key for URL struct
	RequestURLKey ContextKey = "request_url"
)

// GetUser retrieves the user entity from the context.
// Returns nil if the user is not found in the context.
func GetUser(ctx context.Context) *entity.User {
	user, ok := ctx.Value(UserKey).(*entity.User)
	if !ok {
		return nil
	}
	return user
}
