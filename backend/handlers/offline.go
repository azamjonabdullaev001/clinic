package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// decrementStockForOrder reduces product stock_quantity when an order is delivered.
func decrementStockForOrder(order models.Order) {
	for _, item := range order.Items {
		var product models.Product
		if database.DB.First(&product, item.ProductID).Error != nil {
			continue
		}
		var packs int
		if item.UnitType == "piece" {
			if product.QuantityPerPack > 0 {
				packs = (item.Quantity + product.QuantityPerPack - 1) / product.QuantityPerPack
			} else {
				packs = 1
			}
		} else {
			packs = item.Quantity
		}
		newStock := product.StockQuantity - packs
		if newStock < 0 {
			newStock = 0
		}
		database.DB.Model(&models.Product{}).Where("id = ?", item.ProductID).
			Update("stock_quantity", newStock)
	}
}

type OfflineItemInput struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	UnitType  string `json:"unit_type"`
}

type OfflineSaleInput struct {
	Items       []OfflineItemInput `json:"items" binding:"required,min=1"`
	OfflineNote string             `json:"offline_note"`
}

func CreateOfflineSale(c *gin.Context) {
	var input OfflineSaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	tx := database.DB.Begin()

	var workerIDPtr *uint
	if wid, ok := c.Get("workerID"); ok {
		if w, ok := wid.(uint); ok {
			workerIDPtr = &w
		}
	}

	order := models.Order{
		UserID:      nil,
		WorkerID:    workerIDPtr,
		Status:      "delivered",
		Phone:       "offline",
		OrderCode:   generateOrderCode(),
		IsOffline:   true,
		OfflineNote: input.OfflineNote,
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании записи"})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании записи"})
			return
		}
	}

	tx.Commit()

	database.DB.Preload("Items.Product").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	decrementStockForOrder(order)

	c.JSON(http.StatusCreated, order)
}
