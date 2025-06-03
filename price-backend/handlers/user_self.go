package handlers

import (
	"net/http"
	"price-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PasswordChangeRequest struct {
	NewPassword string `json:"new_password" binding:"required"`
}

// ChangePassword Изменение пароля (для менеджера и клиента)
// @Summary Изменить пароль
// @Tags Manager
// @Tags Client
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body PasswordChangeRequest true "Новый пароль"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /manager/change-password [post]
// @Router /client/change-password [post]
func ChangePassword(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, _ := c.Get("user_id")
		var req PasswordChangeRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		hash, err := utils.HashPassword(req.NewPassword)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		if err := db.Exec(`UPDATE users SET password_hash = ? WHERE id = ?`, hash, userID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Password updated"})
	}
}
