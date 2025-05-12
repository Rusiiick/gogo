package models

import "time"

type RefreshToken struct {
	UserID    int       `json:"user_id"`
	Token     string    `json:"refresh_token"`
	ExpiresAt time.Time `json:"expires_at"`
}
