package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	"github.com/arfanxn/welding/internal/module/shared/contextkey"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var _ Middleware = (*authenticateMiddleware)(nil)

type AuthenticateMiddleware interface {
	Middleware
}

type authenticateMiddleware struct {
	UserRepository userRepository.UserRepository
	JWTService     jwt.JWTService
}

type NewAuthenticateMiddlewareParams struct {
	fx.In

	UserRepository userRepository.UserRepository
	JWTService     jwt.JWTService
}

func NewAuthenticateMiddleware(
	params NewAuthenticateMiddlewareParams,
) (AuthenticateMiddleware, error) {
	return &authenticateMiddleware{
		UserRepository: params.UserRepository,
		JWTService:     params.JWTService,
	}, nil
}

// MiddlewareFunc returns a Gin middleware handler function that handles JWT authentication
// and user verification for protected routes.
func (m *authenticateMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Get the Authorization header from the request
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "Header Authorization diperlukan", nil))
		}

		// 2. Extract and validate the Bearer token format
		// Expected format: "Bearer <token>"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "Format header Authorization tidak valid. Format yang benar: Bearer <token>", nil))
		}

		// 3. Verify the JWT token and extract claims
		tokenStr := tokenParts[1]
		claims, err := m.JWTService.VerifyToken(tokenStr)
		if err != nil {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "Token tidak valid atau sudah kadaluarsa", nil))
		}

		// 4. Verify that the user exists in the database
		user, err := m.UserRepository.Find(claims.UserID)
		if err != nil {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "User tidak ditemukan", nil))
		}

		// 5. Check if the user account is active
		if !user.IsActive() {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "User tidak aktif, silahkan hubungi admin", nil))
		}

		// 6. Store user information in both Gin context and request context
		// This makes the user data available to subsequent handlers
		ctx := c.Request.Context()
		for key, value := range map[contextkey.ContextKey]any{
			contextkey.UserIdKey: user.Id, // Store user ID
			contextkey.ClaimsKey: claims,  // Store JWT claims
			contextkey.UserKey:   user,    // Store full user object
		} {
			c.Set(key, value)                        // Set in Gin context
			ctx = context.WithValue(ctx, key, value) // Set in request context
		}
		c.Request = c.Request.WithContext(ctx)

		// 7. Proceed to the next middleware/handler in the chain
		c.Next()
	}
}
