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

func (db *DB) CreateSuppliy(ctx context.Context, supplies models.Supply) error {
	_, err := db.Conn.Exec(ctx, "INSERT INTO public.supplies (title, description, price, quantity) VALUES ($1, $2, $3, $4)", supplies.Title, supplies.Description, supplies.Price, supplies.Quantity)
	if err != nil {
		log.Println("Ошибка при добавлении товара", err)
		return err
	}
	return nil
}

func (db *DB) GetSupplyByID(ctx context.Context, id int) (*models.Supply, error) {
	var supply models.Supply

	err := db.Conn.QueryRow(ctx, "SELECT  id, title, description, price, quantity FROM public.supplies where id = $1", id).Scan(&supply.ID, &supply.Title, &supply.Description, &supply.Price, &supply.Quantity)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("Товар с таким названием не найден")
		} else {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		}
		return nil, err
	}

	return &supply, nil
}

func (db *DB) UpdateSupply(ctx context.Context, supply models.Supply, id int) error {

	cmdTag, err := db.Conn.Exec(ctx, "UPDATE public.supplies SET title = $1, description = $2, price = $3, quantity = $4 WHERE id = $5", supply.Title, supply.Description, supply.Price, supply.Quantity, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при обновлении данных: %v\n", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		fmt.Println("Товар с таким ID не найден")
	} else {
		fmt.Println("Товар обновлен")
	}

	return nil
}

func (db *DB) DeleteSupplyByID(ctx context.Context, id int) error {
	_, err := db.Conn.Exec(ctx, "DELETE FROM public.supplies WHERE id = $1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при удалении товара: %v\n", err)
		return err
	}
	return nil
}
