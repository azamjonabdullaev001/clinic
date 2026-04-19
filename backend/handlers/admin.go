package handlers

import (
	"clinic-backend/database"
	"clinic-backend/middleware"
	"clinic-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminLoginInput struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminSettingsInput struct {
	Phone       string `json:"phone"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func AdminLogin(c *gin.Context) {
	var input AdminLoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Номер телефона и пароль обязательны"})
		return
	}

	var admin models.Admin
	if err := database.DB.Where("phone = ?", input.Phone).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
		return
	}

	token, err := middleware.GenerateToken(admin.ID, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"admin": gin.H{
			"id":    admin.ID,
			"phone": admin.Phone,
		},
	})
}

func AdminProfile(c *gin.Context) {
	adminID, _ := c.Get("adminID")
	var admin models.Admin
	if err := database.DB.First(&admin, adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Администратор не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    admin.ID,
		"phone": admin.Phone,
	})
}

func UpdateAdminSettings(c *gin.Context) {
	adminID, _ := c.Get("adminID")
	var input AdminSettingsInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	var admin models.Admin
	if err := database.DB.First(&admin, adminID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Администратор не найден"})
		return
	}

	if input.NewPassword != "" {
		if input.OldPassword == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Текущий пароль обязателен"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.OldPassword)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный текущий пароль"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(input.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		admin.Password = string(hash)
	}

	if input.Phone != "" && phoneRegex.MatchString(input.Phone) {
		admin.Phone = input.Phone
	}

	if err := database.DB.Save(&admin).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Настройки обновлены",
		"admin": gin.H{
			"id":    admin.ID,
			"phone": admin.Phone,
		},
	})
}
