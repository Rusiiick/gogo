package api

import (
	"candy_shop/internal/models"
	"time"
)

type UserDTO struct {
	ID        int             `json:"id"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	Role      models.UserRole `json:"role"`
	CreatedAt time.Time       `json:"created_at"`
}

func ToDTOuser(u *models.User) *UserDTO {
	if u == nil {
		return nil
	}

	user := &UserDTO{
		ID:        u.ID,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}

	return user
}

func DTOtoModelsUser(u *UserDTO) *models.User {
	return &models.User{
		ID:        u.ID,
		Email:     u.Email,
		Password:  u.Password,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}
