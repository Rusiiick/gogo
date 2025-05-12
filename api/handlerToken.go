package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshHandler(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh токен не найден в cookie"})
		return
	}

	rt, err := h.db.GetRefreshToken(c.Request.Context(), refreshToken)
	if err != nil {
		log.Printf("Ошибка при получении refresh токена: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
		return
	}

	if time.Now().After(rt.ExpiresAt) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Срок действия токена истёк, войдите заново"})
		return
	}

	accessToken, err := h.ts.GenerateJWT(rt.UserID, 0)
	if err != nil {
		log.Printf("Ошибка при генерации access токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	accessTokenMaxAge := int(h.cfg.AccessTokenDuration.Seconds())
	refreshTokenTTL := h.cfg.RefreshTokenDuration

	c.SetCookie("access_token", accessToken, accessTokenMaxAge, "/", "", false, true)

	newRefreshToken, err := h.ts.GenerateRefreshToken()
	if err != nil {
		log.Printf("Ошибка при генерации нового refresh токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка генерации токена"})
		return
	}

	expiresAt := time.Now().Add(refreshTokenTTL)

	err = h.db.UpsertRefreshToken(c.Request.Context(), rt.UserID, newRefreshToken, expiresAt)
	if err != nil {
		log.Printf("Ошибка при сохранении refresh токена: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сохранения токена"})
		return
	}

	c.SetCookie("refresh_token", newRefreshToken, int(refreshTokenTTL.Seconds()), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Токены обновлены"})
}
