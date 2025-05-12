package models

type Category struct {
	ID          int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
