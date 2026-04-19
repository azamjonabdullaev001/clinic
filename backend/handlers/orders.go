package handlers

import (
	"bytes"
	"clinic-backend/config"
	"clinic-backend/database"
	"clinic-backend/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type OrderItemInput struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	UnitType  string `json:"unit_type"`
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

	database.DB.Preload("Items.Product").Preload("User").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	go sendTelegramNotification(order)

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

func formatNumber(n float64) string {
	s := fmt.Sprintf("%.0f", n)
	if len(s) <= 3 {
		return s
	}
	var result strings.Builder
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteRune(' ')
		}
		result.WriteRune(c)
	}
	return result.String()
}

func sendTelegramNotification(order models.Order) {
	cfg := config.Load()
	if cfg.TelegramBotToken == "" || cfg.TelegramChatID == "" {
		return
	}

	var sb strings.Builder
	sb.WriteString("🔔 НОВЫЙ ЗАКАЗ\n\n")

	var totalSum float64
	for _, item := range order.Items {
		item.Product.ComputePackPrice()
		unitLabel := "коробка"
		if item.UnitType == "piece" {
			unitLabel = "шт"
		}
		sb.WriteString(fmt.Sprintf("💊 Лекарство: %s\n", item.Product.Name))
		sb.WriteString(fmt.Sprintf("📦 Количество: %d %s\n", item.Quantity, unitLabel))
		sb.WriteString(fmt.Sprintf("💰 Сумма: %s сум\n\n", formatNumber(item.Price)))
		totalSum += item.Price
	}

	if len(order.Items) > 1 {
		sb.WriteString(fmt.Sprintf("💰 Общая сумма: %s сум\n\n", formatNumber(totalSum)))
	}

	sb.WriteString(fmt.Sprintf("👤 Имя: %s\n", order.User.FirstName))
	sb.WriteString(fmt.Sprintf("👨 Фамилия: %s\n", order.User.LastName))
	if order.User.MiddleName != "" {
		sb.WriteString(fmt.Sprintf("👤 Отчество: %s\n", order.User.MiddleName))
	}
	sb.WriteString(fmt.Sprintf("📱 Телефон: +%s\n", order.Phone))
	sb.WriteString("🌍 Страна: O'zbekiston")

	payload := map[string]string{
		"chat_id": cfg.TelegramChatID,
		"text":    sb.String(),
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Telegram marshal error: %v", err)
		return
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TelegramBotToken)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Telegram send error: %v", err)
		return
	}
	defer resp.Body.Close()
}
