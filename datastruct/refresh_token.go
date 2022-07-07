package datastruct

import (
	"github.com/golang-jwt/jwt/v4"
)

type RefreshTokenPayload struct {
	Generation int16 `json:"generation"`
	jwt.StandardClaims
}
