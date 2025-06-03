package main

import (
	"log"
	"price-backend/config"
	"price-backend/db"
	_ "price-backend/docs"
	"price-backend/routes"

	"github.com/gin-gonic/gin"
)

// @title Price Monitor API
// @version 1.0
// @description REST API для мониторинга цен

// @contact.name SmartStack Dev
// @contact.email admin@costpulse.ru

// @host costpulse.ru
// @BasePath /api

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.LoadConfig()
	database := db.ConnectDB(cfg)

	r := gin.Default()
	routes.RegisterRoutes(r, database, cfg)

	log.Println("🚀 Server running on port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
