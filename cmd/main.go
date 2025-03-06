package main

import (
	"candy_shop/config"
	"candy_shop/internal/api"
	"candy_shop/internal/database"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	config := config.DefaultConfig()

	fmt.Printf("Config: %+v\n", config)

	db, err := database.NewDB(ctx, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(ctx)

	handler := api.NewHandler(db)

	router := gin.Default()

	router.POST("/supplies", handler.CreateSup)

	router.GET("/supplies/:id", handler.GetSupply)

	router.PATCH("/supplies/:id", handler.UpdateSup)

	router.DELETE("/supplies/:id", handler.DeleteSup)

	router.Run("localhost:8080")
}
