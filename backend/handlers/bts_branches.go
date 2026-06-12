package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBtsBranches(c *gin.Context) {
	var branches []models.BtsBranch
	database.DB.Where("is_active = ?", true).Order("region, city, name").Find(&branches)
	c.JSON(http.StatusOK, branches)
}

func CreateBtsBranch(c *gin.Context) {
	var b models.BtsBranch
	if err := c.ShouldBindJSON(&b); err != nil || b.Name == "" || b.City == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заполните название и город"})
		return
	}
	b.IsActive = true
	if err := database.DB.Create(&b).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании"})
		return
	}
	c.JSON(http.StatusCreated, b)
}

func UpdateBtsBranch(c *gin.Context) {
	id := c.Param("id")
	var b models.BtsBranch
	if err := database.DB.First(&b, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Не найдено"})
		return
	}
	var input models.BtsBranch
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	b.Name = input.Name
	b.City = input.City
	b.Region = input.Region
	b.Address = input.Address
	b.Lat = input.Lat
	b.Lng = input.Lng
	b.IsActive = input.IsActive
	database.DB.Save(&b)
	c.JSON(http.StatusOK, b)
}

func DeleteBtsBranch(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.BtsBranch{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Удалено"})
}
