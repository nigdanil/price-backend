package handlers

import (
	"net/http"
	"price-backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetProductPrices Получение истории изменения цен с фильтрами и пагинацией
// @Summary История цен
// @Tags Gallery
// @Security BearerAuth
// @Produce json
// @Param from query string false "Дата от (yyyy-mm-dd)"
// @Param to query string false "Дата до (yyyy-mm-dd)"
// @Param page query int false "Номер страницы" default(1)
// @Param limit query int false "Размер страницы" default(50)
// @Param category_id query string false "ID категории"
// @Success 200 {array} models.GalleryProductPrice
// @Failure 500 {object} map[string]string
// @Router /gallery/prices [get]
func GetProductPrices(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var results []models.GalleryProductPrice

		// Параметры
		from := c.Query("from")
		to := c.Query("to")
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "50")
		categoryID := c.Query("category_id")

		// Парсинг
		pageNum, _ := strconv.Atoi(page)
		limitNum, _ := strconv.Atoi(limit)
		offset := (pageNum - 1) * limitNum

		query := db.Model(&models.GalleryProductPrice{})

		if from != "" {
			if t, err := time.Parse("2006-01-02", from); err == nil {
				query = query.Where("updated_at >= ?", t)
			}
		}
		if to != "" {
			if t, err := time.Parse("2006-01-02", to); err == nil {
				query = query.Where("updated_at <= ?", t)
			}
		}
		if categoryID != "" {
			query = query.Where("category_id = ?", categoryID)
		}

		err := query.Order("updated_at DESC").
			Offset(offset).
			Limit(limitNum).
			Find(&results).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении данных"})
			return
		}

		c.JSON(http.StatusOK, results)
	}
}
