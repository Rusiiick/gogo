package api

import (
	"candy_shop/internal/models"
)

type CategoryDTO struct {
	ID          int    `json:"category_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToDTOctgry(c *models.Category) *CategoryDTO {
	if c == nil {
		return nil
	}

	category := &CategoryDTO{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}

	return category
}

func DTOtoModelsctgry(c *CategoryDTO) *models.Category {
	return &models.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
