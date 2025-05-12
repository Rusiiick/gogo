package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) RegisterHandler(c *gin.Context) {
	var user UserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("Ошибка при разборе запроса: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Ошибка при регистрации пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при регистрации"})
		return
	}

	user.Password = string(passwordHash)

	userID, err := h.db.AddUser(c.Request.Context(), DTOtoModelsUser(&user))
	if err != nil {
		log.Printf("Ошибка при регистрации пользователя: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при регистрации"})
		return
	}

	log.Printf("✅ Пользователь зарегистрирован: ID=%d", userID)
	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь зарегистрирован",
		"user_id": userID,
	})
}

func (h *Handler) LoginHandler(c *gin.Context) {
	var userDTO UserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		log.Printf("Ошибка при разборе запроса: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.db.GetUserByEmail(c.Request.Context(), userDTO.Email)
	if err != nil {
		log.Printf("Ошибка при получении пользователя из БД: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверная почта или пароль"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDTO.Password))
	if err != nil {
		log.Printf("Пароль не совпадает: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверная почта или пароль"})
		return
	}

	tokenString, err := h.ts.GenerateJWT(user.ID, 0)
	if err != nil {
		log.Printf("Ошибка при создании токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании токена"})
		return
	}

	accessTokenMaxAge := int(h.cfg.AccessTokenDuration.Seconds())
	refreshTokenTTL := h.cfg.RefreshTokenDuration

	c.SetCookie("access_token", tokenString, accessTokenMaxAge, "/", "", false, true)

	refreshToken, err := h.ts.GenerateRefreshToken()
	if err != nil {
		log.Printf("Ошибка при создании refresh токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании токена"})
		return
	}

	expiresAt := time.Now().Add(refreshTokenTTL)

	err = h.db.UpsertRefreshToken(c.Request.Context(), user.ID, refreshToken, expiresAt)
	if err != nil {
		log.Printf("Ошибка при сохранении refresh токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении токена"})
		return
	}

	c.SetCookie("refresh_token", refreshToken, int(refreshTokenTTL.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Успешный вход в систему",
		"user_id": user.ID,
	})
}
