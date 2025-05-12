package auth

import (
	"candy_shop/internal/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	secret              []byte
	accessTokenDuration time.Duration
}

func NewTokenService(secret string, accessTokenDuration time.Duration) *TokenService {
	return &TokenService{
		secret:              []byte(secret),
		accessTokenDuration: accessTokenDuration,
	}
}

func (ts *TokenService) GenerateJWT(userID int, role int8) (string, error) {
	claims := models.Claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ts.accessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(ts.secret)
}
