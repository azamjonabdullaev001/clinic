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
		if item.Quantity <= 0 {
			continue
		}
		var product models.Product
		if database.DB.First(&product, item.ProductID).Error != nil {
			continue
		}
		pieces := pieceCount(product, item.Quantity, item.UnitType)
		newStock := product.StockQuantity - pieces
		if newStock < 0 {
			newStock = 0
		}
		database.DB.Model(&models.Product{}).Where("id = ?", item.ProductID).
			Update("stock_quantity", newStock)
		BroadcastStock(item.ProductID, newStock)
	}
}

type OfflineItemInput struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	UnitType  string `json:"unit_type"`
}

// validOfflinePayment reports whether m is an accepted offline payment method.
func validOfflinePayment(m string) bool {
	return m == "cash" || m == "terminal" || m == "card"
}

type OfflineSaleInput struct {
	Items         []OfflineItemInput `json:"items" binding:"required,min=1"`
	OfflineNote   string             `json:"offline_note"`
	IsVIP         bool               `json:"is_vip"`
	PaymentMethod string             `json:"payment_method"` // "cash", "terminal", "card"
	CardType      string             `json:"card_type"`      // "humo", "uzcard", "visa", "mastercard"
	ReferredBy    string             `json:"referred_by"`    // doctor who referred the patient
	SalesChannel  string             `json:"sales_channel"`  // marketplace for manager sales
}

func CreateOfflineSale(c *gin.Context) {
	var input OfflineSaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// Determine payment method. Manager marketplace sales are paid through the platform;
	// "own patient" sales are free; otherwise the cashier picks cash/terminal/card.
	paymentMethod := input.PaymentMethod
	cardType := ""
	if input.SalesChannel != "" {
		paymentMethod = "marketplace"
	} else if input.IsVIP {
		paymentMethod = ""
	} else {
		if paymentMethod == "" {
			paymentMethod = "cash"
		}
		if !validOfflinePayment(paymentMethod) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный способ оплаты"})
			return
		}
	}

	tx := database.DB.Begin()

	var workerIDPtr *uint
	if wid, ok := c.Get("workerID"); ok {
		if w, ok := wid.(uint); ok {
			workerIDPtr = &w
		}
	}

	// Offline & marketolog sales draw from the shared warehouse and cannot exceed it.
	if err := reserveProductStock(tx, input.Items); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		UserID:        nil,
		WorkerID:      workerIDPtr,
		Status:        "delivered",
		Phone:         "offline",
		OrderCode:     generateNurseCode(), // offline orders use a 5-digit code
		IsOffline:     true,
		IsVIP:         input.IsVIP,
		PaymentMethod: paymentMethod,
		CardType:      cardType,
		SalesChannel:  input.SalesChannel,
		ReferredBy:    input.ReferredBy,
		OfflineNote:   input.OfflineNote,
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
		if unitType != "piece" {
			unitType = "pack"
		}
		var price float64
		if !input.IsVIP {
			if unitType == "piece" {
				price = product.PricePerPill * float64(item.Quantity)
			} else {
				price = product.PricePerPack * float64(item.Quantity)
			}
		}

		orderItem := models.OrderItem{
			OrderID:          order.ID,
			ProductID:        item.ProductID,
			Quantity:         item.Quantity,
			OriginalQuantity: item.Quantity,
			UnitType:         unitType,
			Price:            price,
		}

		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании записи"})
			return
		}
	}

	tx.Commit()

	for _, item := range input.Items {
		broadcastProductStock(item.ProductID)
	}

	database.DB.Preload("Items.Product").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusCreated, order)
}
