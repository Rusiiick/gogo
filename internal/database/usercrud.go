package database

import (
	"candy_shop/internal/models"
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func (db *DB) AddUser(ctx context.Context, user *models.User) (int, error) {
	var userID int

	err := db.Conn.QueryRow(ctx, "INSERT INTO public.users (email, password_hash, role, created_at) VALUES ($1, $2, $3, NOW()) RETURNING id", user.Email, user.Password, models.RoleUser).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func (db *DB) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}

	err := db.Conn.QueryRow(ctx, "SELECT id, email, password_hash, role, created_at FROM public.users WHERE email = $1", email).Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("User с таким email не найден")
			return nil, nil
		}

		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}

	return user, nil
}
