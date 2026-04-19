package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderItemInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type CreateOrderInput struct {
	Items []OrderItemInput `json:"items" binding:"required,min=1"`
	Phone string           `json:"phone" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	userID, _ := c.Get("userID")

	var input CreateOrderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные заказа"})
		return
	}

	order := models.Order{
		UserID: userID.(uint),
		Status: "pending",
		Phone:  input.Phone,
	}

	tx := database.DB.Begin()

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
		return
	}

	for _, item := range input.Items {
		var product models.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Препарат не найден"})
			return
		}

		product.ComputePackPrice()
		orderItem := models.OrderItem{
			OrderID:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.PricePerPack * float64(item.Quantity),
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании заказа"})
			return
		}
	}

	tx.Commit()

	database.DB.Preload("Items.Product").Preload("User").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusCreated, order)
}

func GetUserOrders(c *gin.Context) {
	userID, _ := c.Get("userID")
	var orders []models.Order
	database.DB.Where("user_id = ?", userID).
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

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Preload("Items.Product").Preload("User").
		Order("created_at desc").
		Find(&orders)

	for i := range orders {
		for j := range orders[i].Items {
			orders[i].Items[j].Product.ComputePackPrice()
		}
	}

	c.JSON(http.StatusOK, orders)
}

func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	var input struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный статус"})
		return
	}

	validStatuses := map[string]bool{
		"pending":   true,
		"confirmed": true,
		"shipped":   true,
		"delivered": true,
		"cancelled": true,
	}
	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный статус"})
		return
	}

	order.Status = input.Status
	database.DB.Save(&order)

	database.DB.Preload("Items.Product").Preload("User").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusOK, order)
}
