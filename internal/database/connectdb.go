package database

import (
	"candy_shop/config"
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	Conn *pgx.Conn
}

func NewDB(ctx context.Context, cfg config.DBConfig) (*DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName,
	)

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка подключения к базе данных: %v\n", err)
		return nil, err
	}

	return &DB{Conn: conn}, nil
}

func (db *DB) Close(ctx context.Context) {
	db.Conn.Close(ctx)
}
