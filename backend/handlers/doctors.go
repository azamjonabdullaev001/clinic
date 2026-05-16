package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDoctors(c *gin.Context) {
	var doctors []models.Doctor
	database.DB.Order("name asc").Find(&doctors)
	c.JSON(http.StatusOK, doctors)
}

func CreateDoctor(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя доктора обязательно"})
		return
	}

	doctor := models.Doctor{Name: input.Name}
	if err := database.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении доктора"})
		return
	}
	c.JSON(http.StatusCreated, doctor)
}

func DeleteDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	database.DB.Delete(&models.Doctor{}, id)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
