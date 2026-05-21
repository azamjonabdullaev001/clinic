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

	// Check admin first
	var admin models.Admin
	if err := database.DB.Where("phone = ?", input.Phone).First(&admin).Error; err == nil {
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
			"role":  "admin",
			"admin": gin.H{"id": admin.ID, "phone": admin.Phone},
		})
		return
	}

	// Check worker
	var worker models.Worker
	if err := database.DB.Where("phone = ?", input.Phone).First(&worker).Error; err == nil {
		if err := bcrypt.CompareHashAndPassword([]byte(worker.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
			return
		}
		token, err := middleware.GenerateToken(worker.ID, "worker")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"token":  token,
			"role":   "worker",
			"worker": gin.H{"id": worker.ID, "phone": worker.Phone, "name": worker.Name, "role": worker.Role},
		})
		return
	}

	// Check doctor
	var doctor models.Doctor
	if err := database.DB.Where("phone = ?", input.Phone).First(&doctor).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
		return
	}

	if doctor.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учетные данные"})
		return
	}

	token, err := middleware.GenerateToken(doctor.ID, "doctor")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":  token,
		"role":   "doctor",
		"doctor": gin.H{"id": doctor.ID, "phone": doctor.Phone, "name": doctor.Name, "specialty": doctor.Specialty},
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

// ==================== Worker Management ====================

type CreateWorkerInput struct {
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

func GetWorkers(c *gin.Context) {
	var workers []models.Worker
	database.DB.Find(&workers)
	result := make([]gin.H, len(workers))
	for i, w := range workers {
		result[i] = gin.H{"id": w.ID, "name": w.Name, "phone": w.Phone, "role": w.Role}
	}
	c.JSON(http.StatusOK, result)
}

func CreateWorker(c *gin.Context) {
	var input CreateWorkerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя, телефон и пароль обязательны (мин. 6 символов)"})
		return
	}

	if !phoneRegex.MatchString(input.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат номера телефона (998XXXXXXXXX)"})
		return
	}

	var existing models.Worker
	if err := database.DB.Where("phone = ?", input.Phone).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Этот номер уже используется"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	role := input.Role
	if role != "nurse" {
		role = "pickup"
	}

	worker := models.Worker{
		Name:     input.Name,
		Phone:    input.Phone,
		Password: string(hash),
		Role:     role,
	}
	if err := database.DB.Create(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании работника"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": worker.ID, "name": worker.Name, "phone": worker.Phone, "role": worker.Role})
}

type UpdateWorkerInput struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func UpdateWorker(c *gin.Context) {
	id := c.Param("id")
	var worker models.Worker
	if err := database.DB.First(&worker, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Работник не найден"})
		return
	}

	var input UpdateWorkerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if input.Name != "" {
		worker.Name = input.Name
	}
	if input.Phone != "" && phoneRegex.MatchString(input.Phone) {
		worker.Phone = input.Phone
	}
	if input.Role == "nurse" || input.Role == "pickup" {
		worker.Role = input.Role
	}
	if len(input.Password) >= 6 {
		hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			return
		}
		worker.Password = string(hash)
	}

	if err := database.DB.Save(&worker).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": worker.ID, "name": worker.Name, "phone": worker.Phone, "role": worker.Role})
}

func DeleteWorker(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Worker{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Работник удалён"})
}
