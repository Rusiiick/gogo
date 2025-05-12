package api

import "candy_shop/internal/models"

type SupplyDTO struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int     `json:"quantity"`
	CategoryID  int     `json:"category_id"`
}

func ToDTO(t *models.Supply) *SupplyDTO {
	if t == nil {
		return nil
	}

	supplies := &SupplyDTO{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Price:       t.Price,
		Quantity:    t.Quantity,
		CategoryID:  t.CategoryID,
	}

	return supplies
}

func DTOtoModels(t *SupplyDTO) *models.Supply {
	return &models.Supply{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Price:       t.Price,
		Quantity:    t.Quantity,
		CategoryID:  t.CategoryID,
	}
}
