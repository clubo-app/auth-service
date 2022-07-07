package service

import (
	"time"

	"github.com/clubo-app/auth-service/datastruct"
	"github.com/clubo-app/auth-service/repository"
	"github.com/clubo-app/packages/types"
	"github.com/clubo-app/packages/utils/middleware"
	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TokenManager interface {
	NewAccessToken(u repository.Account) (string, error)
	ValidateAccessToken(token string) (types.AccessTokenPayload, error)
	NewRefreshToken(id string, generation int16) (string, error)
	ValidateRefreshToken(tokensStr string) (datastruct.RefreshTokenPayload, error)
}

type tokenManager struct {
	secret string
}

const (
	ACCESS_TOKEN_EXP = time.Hour * 24 * 14
)

func NewTokenManager(secret string) TokenManager {
	return tokenManager{secret: secret}
}

func (t tokenManager) ValidateAccessToken(tokensStr string) (types.AccessTokenPayload, error) {
	token, err := jwt.Parse(tokensStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Error(codes.Unauthenticated, "Unexpected signing method")
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(t.secret), nil
	})
	if err != nil {
		return types.AccessTokenPayload{}, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return middleware.ParseAccessTokenMapClaims(claims), nil
	}

	return types.AccessTokenPayload{}, nil
}

func (t tokenManager) ValidateRefreshToken(tokensStr string) (datastruct.RefreshTokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokensStr, &datastruct.RefreshTokenPayload{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, status.Error(codes.Unauthenticated, "Unexpected signing method")
		}

		return []byte(t.secret), nil
	})
	if err != nil {
		return datastruct.RefreshTokenPayload{}, nil
	}

	if claims, ok := token.Claims.(*datastruct.RefreshTokenPayload); ok && token.Valid {
		return *claims, nil
	}

	return datastruct.RefreshTokenPayload{}, nil
}

func (t tokenManager) NewAccessToken(u repository.Account) (string, error) {
	claims := jwt.MapClaims{
		"sub":           u.ID,
		"iss":           "clubo-app.com",
		"emailVerified": u.EmailVerified,
		"role":          u.Type,
		"iat":           time.Now().Unix(),
		"exp":           time.Now().AddDate(0, 0, 14).Unix(),
		"provider":      u.Provider.ToGRPCProvider().String(),
	}

	if u.Provider.Valid {
		val, err := u.Provider.Value()
		if err != nil {
			return "", err
		}
		claims["provider"] = val
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}

func (t tokenManager) NewRefreshToken(id string, generation int16) (string, error) {
	claims := datastruct.RefreshTokenPayload{
		generation,
		jwt.StandardClaims{
			Subject:   id,
			Issuer:    "clubo-app.com",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(t.secret))
}
