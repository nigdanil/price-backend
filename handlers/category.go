package handlers

import (
	"net/http"
	"price-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCategoryURLs Получение всех URL категорий
// @Summary Получение всех URL категорий
// @Tags Categories
// @Security BearerAuth
// @Produce json
// @Success 200 {array} models.CategoryURL
// @Failure 500 {object} map[string]string
// @Router /categories/urls [get]
func GetCategoryURLs(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var urls []models.CategoryURL
		if err := db.Find(&urls).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch category URLs"})
			return
		}
		c.JSON(http.StatusOK, urls)
	}
}
