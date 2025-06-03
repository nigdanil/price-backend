package handlers

import (
	"net/http"
	"price-backend/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProductByID Получение продукта по ID
// @Summary Получить продукт по ID
// @Tags Gallery
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID продукта"
// @Success 200 {object} models.GalleryProduct
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /gallery/products/{id} [get]
func GetProductByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var product models.GalleryProduct
		if err := db.First(&product, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusOK, product)
	}
}

// GetProducts Получение списка продуктов с пагинацией
// @Summary Список продуктов
// @Tags Gallery
// @Security BearerAuth
// @Produce json
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Размер страницы" default(50)
// @Success 200 {object} models.PaginatedProductsResponse
// @Failure 500 {object} map[string]string
// @Router /gallery/products [get]
func GetProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "50")

		page, _ := strconv.Atoi(pageStr)
		limit, _ := strconv.Atoi(limitStr)

		if limit <= 0 {
			limit = 50
		} else if limit > 200 {
			limit = 200
		}

		offset := (page - 1) * limit

		var products []models.GalleryProduct
		var total int64

		if err := db.Model(&models.GalleryProduct{}).Count(&total).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
			return
		}

		if err := db.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
			return
		}

		c.JSON(http.StatusOK, models.PaginatedProductsResponse{
			Products: products,
			Page:     page,
			Limit:    limit,
			Total:    total,
		})
	}
}

// GetCategories Получение списка категорий
// @Summary Получение списка категорий
// @Tags Gallery
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.GalleryCategory
// @Failure 500 {object} map[string]string
// @Router /gallery/categories [get]
func GetCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.GalleryCategory
		if err := db.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
			return
		}
		c.JSON(http.StatusOK, categories)
	}
}
