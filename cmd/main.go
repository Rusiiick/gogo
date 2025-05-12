package main

import (
	"candy_shop/api"
	"candy_shop/config"
	"candy_shop/internal/auth"
	"candy_shop/internal/database"
	"candy_shop/internal/middleware"

	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	config, err := config.LoadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse config: %v\n", err)
		os.Exit(1)
	}

	db, err := database.NewDB(ctx, config.DBconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(ctx)

	ts := auth.NewTokenService(config.JWTSecret, config.AccessTokenDuration)

	handler := api.NewHandler(db, ts, config)

	router := gin.Default()

	router.Use(middleware.LoggerMiddleware())

	router.POST("/user", handler.RegisterHandler)
	router.POST("/user/login", handler.LoginHandler)

	auth := router.Group("/", middleware.AuthMiddleware(ts))
	{
		auth.GET("/supplies/:id", handler.GetSupply)
		auth.GET("/categories", handler.GetAllCtgry)
		auth.GET("/category/:id", handler.GetCtgryByID)
		auth.GET("/supplies/category/:id", handler.GetSupplyByCtgry)
		auth.POST("/user/refresh", handler.RefreshHandler)
	}

	admin := router.Group("/", middleware.AuthMiddleware(ts), middleware.AuthorizeRole(1))
	{
		admin.POST("/supplies", handler.CreateSup)
		admin.PATCH("/supplies/:id", handler.UpdateSup)
		admin.DELETE("/supplies/:id", handler.DeleteSup)

		admin.POST("/category", handler.CreateCtgry)
		admin.PATCH("/category/:id", handler.UpdateCtgry)
		admin.DELETE("/category/:id", handler.DeleteCtgry)
	}

	router.Run("localhost:8080")
}
