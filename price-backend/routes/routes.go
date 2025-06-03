package routes

import (
	"price-backend/config"
	"price-backend/handlers"
	"price-backend/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB, cfg config.Config) {
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.AuthLogger(db))

	// ✅ Swagger route (должен идти первым — без middleware)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	// Admin
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(cfg), middleware.AdminOnly())
	admin.DELETE("/users/:id", handlers.DeleteUser(db))

	// Manager
	manager := api.Group("/manager")
	manager.Use(middleware.AuthMiddleware(cfg), middleware.ManagerOnly())
	manager.POST("/change-password", handlers.ChangePassword(db))

	// Client
	client := api.Group("/client")
	client.Use(middleware.AuthMiddleware(cfg), middleware.ClientOnly())
	client.POST("/change-password", handlers.ChangePassword(db))

	// Auth
	auth := api.Group("/auth")
	{
		auth.POST("/login", handlers.LoginHandler(cfg, db))
		// auth.POST("/register", middleware.AuthMiddleware(cfg), middleware.AdminOnly(), handlers.RegisterHandler(cfg, db))
		auth.POST("/register", handlers.RegisterHandler(cfg, db))

	}

	// Gallery
	gallery := api.Group("/gallery")
	{
		gallery.Use(middleware.AuthMiddleware(cfg))

		gallery.GET("/products/:id", handlers.GetProductByID(db))
		gallery.GET("/products", handlers.GetProducts(db))
		gallery.GET("/categories", handlers.GetCategories(db))
		gallery.GET("/prices", handlers.GetProductPrices(db))

	}

	// Category URLs
	category := api.Group("/categories")
	{
		category.Use(middleware.AuthMiddleware(cfg))
		category.GET("/urls", handlers.GetCategoryURLs(db))
	}
}
