package api

import "candy_shop/category"

type CategoryDTO struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func ToDTOctgry(c *category.Category) *CategoryDTO {
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

func DTOtoModelsctgry(c *CategoryDTO) *category.Category {
	return &category.Category{
		ID:          c.ID,
		Name:        c.Name,
		Description: c.Description,
	}
}
