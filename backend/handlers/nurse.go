package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

func generateNurseCode() string {
	for {
		code := fmt.Sprintf("%05d", rand.Intn(90000)+10000)
		var count int64
		database.DB.Model(&models.Order{}).Where("order_code = ?", code).Count(&count)
		if count == 0 {
			return code
		}
	}
}

type CreateNurseOrderInput struct {
	PatientFirstName string             `json:"patient_first_name" binding:"required"`
	PatientLastName  string             `json:"patient_last_name" binding:"required"`
	Items            []OfflineItemInput `json:"items" binding:"required,min=1"`
}

func CreateNurseOrder(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	wid := workerID.(uint)

	var input CreateNurseOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя, фамилия пациента и препараты обязательны"})
		return
	}

	tx := database.DB.Begin()

	order := models.Order{
		UserID:       nil,
		WorkerID:     &wid,
		Status:       "pending",
		Phone:        "offline",
		OrderCode:    generateNurseCode(),
		IsOffline:    true,
		IsNurseOrder: true,
		PatientFName: input.PatientFirstName,
		PatientLName: input.PatientLastName,
		OfflineNote:  input.PatientFirstName + " " + input.PatientLastName,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
		return
	}

	for _, item := range input.Items {
		var product models.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Товар не найден"})
			return
		}
		product.ComputePackPrice()

		unitType := item.UnitType
		if unitType == "" {
			unitType = "pack"
		}

		var price float64
		if unitType == "piece" {
			price = product.PricePerPill * float64(item.Quantity)
		} else {
			price = product.PricePerPack * float64(item.Quantity)
		}

		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			UnitType:  unitType,
			Price:     price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
			return
		}
	}

	tx.Commit()

	database.DB.Preload("Items.Product").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusCreated, order)
}

func GetNurseOrders(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	var orders []models.Order
	database.DB.Where("worker_id = ? AND is_nurse_order = true", workerID).
		Preload("Items.Product").
		Order("created_at desc").
		Find(&orders)

	for i := range orders {
		for j := range orders[i].Items {
			orders[i].Items[j].Product.ComputePackPrice()
		}
	}

	c.JSON(http.StatusOK, orders)
}

func ConfirmNurseOrder(c *gin.Context) {
	code := c.Param("code")
	var order models.Order
	if err := database.DB.Preload("Items.Product").
		Where("order_code = ? AND is_nurse_order = true", code).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusOK, order)
}
