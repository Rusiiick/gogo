package api

import (
	"candy_shop/config"
	"candy_shop/internal/auth"
	"candy_shop/internal/database"
)

type Handler struct {
	db  *database.DB
	ts  *auth.TokenService
	cfg *config.Config
}

func NewHandler(db *database.DB, ts *auth.TokenService, cfg *config.Config) *Handler {
	return &Handler{
		db:  db,
		ts:  ts,
		cfg: cfg,
	}
}
