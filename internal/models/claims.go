package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID int  `json:"user_id"`
	Role   int8 `json:"role"`
	jwt.RegisteredClaims
}
