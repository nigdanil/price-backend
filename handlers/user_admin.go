package handlers

import (
	"net/http"
	"price-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteUser Удаление пользователя (только для админа)
// @Summary Удаление пользователя
// @Tags Admin
// @Security BearerAuth
// @Param id path string true "ID пользователя"
// @Success 200 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /admin/users/{id} [delete]
func DeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.First(&user, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		if user.Role == models.RoleAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Cannot delete admin users"})
			return
		}

		if err := db.Delete(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
	}
}
