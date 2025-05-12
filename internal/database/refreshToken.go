package database

import (
	"candy_shop/internal/models"
	"context"
	"log"
	"time"
)

func (db *DB) UpsertRefreshToken(ctx context.Context, userID int, refreshToken string, expiresAt time.Time) error {
	query := "INSERT INTO refresh_tokens (user_id, refresh_token, expires_at) VALUES ($1, $2, $3) ON CONFLICT (user_id) DO UPDATE SET refresh_token = EXCLUDED.refresh_token, expires_at = EXCLUDED.expires_at"
	_, err := db.Conn.Exec(ctx, query, userID, refreshToken, expiresAt)
	if err != nil {
		log.Printf("Ошибка при сохранении refresh токена: %v", err)
		return err
	}
	return nil
}

func (db *DB) GetRefreshToken(ctx context.Context, token string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	query := "SELECT user_id, refresh_token, expires_at FROM refresh_tokens WHERE refresh_token = $1"
	err := db.Conn.QueryRow(ctx, query, token).Scan(
		&rt.UserID,
		&rt.Token,
		&rt.ExpiresAt,
	)
	if err != nil {
		log.Printf("Ошибка при получении refresh токена: %v", err)
		return nil, err
	}

	return &rt, nil
}
