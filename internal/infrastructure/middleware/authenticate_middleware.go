package middleware

import (
	"net/http"
	"strings"

	"github.com/arfanxn/welding/internal/infrastructure/http/jwt"
	userRepository "github.com/arfanxn/welding/internal/module/user/domain/repository"
	"github.com/arfanxn/welding/pkg/errorutil"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var _ Middleware = (*AuthenticateMiddleware)(nil)

type AuthenticateMiddleware struct {
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
) (*AuthenticateMiddleware, error) {
	return &AuthenticateMiddleware{
		UserRepository: params.UserRepository,
		JWTService:     params.JWTService,
	}, nil
}

func (m *AuthenticateMiddleware) MiddlewareFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "Header Authorization diperlukan", nil))
		}

		// Extract the token from the Authorization header (format: "Bearer <token>")
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "Format header Authorization tidak valid. Format yang benar: Bearer <token>", nil))
		}

		tokenStr := tokenParts[1]
		claims, err := m.JWTService.VerifyToken(tokenStr)
		if err != nil {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "Token tidak valid atau sudah kadaluarsa", nil))
		}

		user, err := m.UserRepository.Find(claims.UserID)
		if err != nil {
			panic(errorutil.NewHttpError(http.StatusUnauthorized, "User tidak ditemukan", nil))
		}

		c.Set("claims", claims)
		c.Set("user", user)
		c.Next()
	}
}
