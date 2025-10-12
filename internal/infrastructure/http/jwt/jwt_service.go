package jwt

import (
	"time"

	"github.com/arfanxn/welding/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type JWTService interface {
	CreateToken(userID string) (string, error)
	VerifyToken(tokenStr string) (*Claims, error)
}

type jwtService struct {
	Duration  time.Duration
	SecretKey string
}

func NewJWTServiceFromConfig(cfg *config.Config) JWTService {
	return &jwtService{
		Duration:  time.Duration(cfg.JWTDuration) * time.Hour,
		SecretKey: cfg.JWTSecret,
	}
}

func (s *jwtService) CreateToken(userID string) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.Duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.SecretKey))
}

func (s *jwtService) VerifyToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr, &Claims{},
		func(t *jwt.Token) (any, error) {
			return []byte(s.SecretKey), nil
		})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}
