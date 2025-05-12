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

func (db *DB) CreateCategory(ctx context.Context, category *models.Category) (int, error) {
	if category == nil {
		return 0, errors.New("передана пустая категория")
	}

	if category.Name == "" || category.Description == "" {
		return 0, errors.New("название и описание категории не могут быть пустыми")
	}

	var categoryID int
	err := db.Conn.QueryRow(
		ctx,
		"INSERT INTO public.category (category_name, category_description) VALUES ($1, $2) RETURNING category_id",
		category.Name, category.Description,
	).Scan(&categoryID)

	if err != nil {
		log.Println("Ошибка при добавлении категории:", err)
		return 0, err
	}

	return categoryID, nil
}
func (db *DB) GetAllCategory(ctx context.Context) ([]*models.Category, error) {
	var ctgry []*models.Category

	rows, err := db.Conn.Query(ctx, "SELECT category_id, category_name, category_description FROM public.category")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		ctg := &models.Category{}
		err := rows.Scan(&ctg.ID, &ctg.Name, &ctg.Description)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при сканировании строки: %v\n", err)
			return nil, err
		}
		ctgry = append(ctgry, ctg)
	}

	return ctgry, nil
}

func (db *DB) GetCategoryById(ctx context.Context, id int) (*models.Category, error) {
	category := &models.Category{}

	err := db.Conn.QueryRow(ctx, "SELECT category_id, category_name, category_description FROM public.category WHERE category_id = $1", id).
		Scan(&category.ID, &category.Name, &category.Description)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			log.Println("Описание с таким названием не найден")
			return nil, nil
		}

		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	return category, nil
}

func (db *DB) GetSupplyByCategoryID(ctx context.Context, id int) ([]*models.Supply, error) {
	var supplies []*models.Supply

	rows, err := db.Conn.Query(ctx, "SELECT id, title, description, price, quantity, category_id FROM public.supplies where category_id = $1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		return nil, err
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	for rows.Next() {
		item := &models.Supply{}
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.Quantity, &item.CategoryID)

		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
			return nil, err
		}

		supplies = append(supplies, item)

	}
	return supplies, nil
}

func (db *DB) UpdateCategory(ctx context.Context, category models.Category, id int) error {
	cmdTag, err := db.Conn.Exec(ctx, "UPDATE public.category SET category_name = $1, category_description = $2 WHERE category_id = $3", category.Name, category.Description, id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при обновлении данных: %v\n", err)
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		fmt.Println("Категория с таким ID не найдена")
	} else {
		fmt.Println("Ктегория обновлена")
	}

	return nil
}

func (db *DB) DeleteCategory(ctx context.Context, id int) error {
	_, err := db.Conn.Exec(ctx, "DELETE FROM public.category WHERE category_id = $1", id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка при удалении категории: %v\n", err)
		return err
	}
	return nil
}
