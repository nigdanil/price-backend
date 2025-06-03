package handlers

import (
	"net/http"
	"price-backend/config"
	"price-backend/models"
	"price-backend/security"
	"price-backend/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"` // 'admin' или 'manager'
}

// LoginHandler Авторизация пользователя
// @Summary Авторизация пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body AuthRequest true "Данные пользователя"
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 429 {object} map[string]interface{}
// @Router /auth/login [post]
func LoginHandler(cfg config.Config, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AuthRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		ip := c.ClientIP()

		// ✅ Проверка блокировки IP
		blocked, retry := security.CheckAndRecordLoginAttempt(ip)
		if blocked {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error":       "Too many attempts. Try again later",
				"retry_after": retry.Seconds(),
			})
			return
		}

		var user models.User
		if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		// ✅ Успешный вход — сброс счетчика
		security.ClearLoginAttempts(ip)

		token, err := utils.GenerateJWT(user.ID.String(), user.Role, cfg.JWTSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

// RegisterHandler Регистрация нового пользователя (доступно только админу)
// @Summary Регистрация нового пользователя
// @Tags Auth
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param register body RegisterRequest true "Данные нового пользователя"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 409 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /auth/register [post]
func RegisterHandler(cfg config.Config, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		if req.Role != "admin" && req.Role != "manager" && req.Role != "client" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
			return
		}

		hash, err := utils.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		user := models.User{
			Username:     req.Username,
			PasswordHash: hash,
			Role:         req.Role,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User registered"})
	}
}
