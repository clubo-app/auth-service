package service

import (
	"time"

	"github.com/clubo-app/auth-service/repository"
	"github.com/golang-jwt/jwt"
)

type TokenManager interface {
	NewJWT(u repository.Account) (string, error)
}

type tokenManager struct {
	secret string
}

func NewTokenManager(secret string) TokenManager {
	return tokenManager{secret: secret}
}

func (t tokenManager) NewJWT(u repository.Account) (string, error) {
	claims := jwt.MapClaims{
		"sub":           u.ID,
		"iss":           "sessions.com",
		"emailVerified": u.EmailVerified,
		"role":          u.Role,
		"iat":           time.Now().Unix(),
	}

	if u.Provider != "" {
		claims["provider"] = u.Provider
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}
