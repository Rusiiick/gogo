package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		log.Printf("📥 Запрос: %s %s", c.Request.Method, c.Request.URL.Path)

		c.Next()

		duration := time.Since(start)
		responseTime := time.Now().Format("2006-01-02 15:04:05")
		log.Printf("📤 [%s] Ответ: %d | %s %s | Время обработки: %v\n", responseTime, c.Writer.Status(), c.Request.Method, c.Request.URL.Path, duration)
	}
}
