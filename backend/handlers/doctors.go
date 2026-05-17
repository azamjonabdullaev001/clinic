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
		Name      string `json:"name" binding:"required"`
		Specialty string `json:"specialty"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя доктора обязательно"})
		return
	}

	doctor := models.Doctor{Name: input.Name, Specialty: input.Specialty}
	if err := database.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении доктора"})
		return
	}
	c.JSON(http.StatusCreated, doctor)
}

func UpdateDoctor(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var doctor models.Doctor
	if err := database.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}
	var input struct {
		Name      string `json:"name" binding:"required"`
		Specialty string `json:"specialty"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя доктора обязательно"})
		return
	}
	doctor.Name = input.Name
	doctor.Specialty = input.Specialty
	database.DB.Save(&doctor)
	c.JSON(http.StatusOK, doctor)
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

type DoctorStatProduct struct {
	ProductID   uint    `json:"product_id"`
	ProductName string  `json:"product_name"`
	TotalPacks  int     `json:"total_packs"`
	TotalPieces int     `json:"total_pieces"`
	OrderCount  int     `json:"order_count"`
	Revenue     float64 `json:"revenue"`
}

func GetDoctorStats(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var doctor models.Doctor
	if err := database.DB.First(&doctor, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}

	var orders []models.Order
	database.DB.Where("referred_by LIKE ?", "%"+doctor.Name+"%").
		Preload("Items.Product").
		Find(&orders)

	statsMap := make(map[uint]*DoctorStatProduct)
	for _, order := range orders {
		for _, item := range order.Items {
			s, ok := statsMap[item.ProductID]
			if !ok {
				s = &DoctorStatProduct{
					ProductID:   item.ProductID,
					ProductName: item.Product.Name,
				}
				statsMap[item.ProductID] = s
			}
			s.Revenue += item.Price
			s.OrderCount++
			if item.UnitType == "piece" {
				s.TotalPieces += item.Quantity
			} else {
				s.TotalPacks += item.Quantity
			}
		}
	}

	var result []DoctorStatProduct
	for _, s := range statsMap {
		result = append(result, *s)
	}
	// Sort by revenue
	for i := 0; i < len(result)-1; i++ {
		for j := i + 1; j < len(result); j++ {
			if result[j].Revenue > result[i].Revenue {
				result[i], result[j] = result[j], result[i]
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"doctor":        doctor,
		"total_orders":  len(orders),
		"products":      result,
	})
}
