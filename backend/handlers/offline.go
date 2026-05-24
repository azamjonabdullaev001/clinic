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

// validOfflinePayment reports whether m is an accepted offline payment method.
func validOfflinePayment(m string) bool {
	return m == "cash" || m == "terminal" || m == "card"
}

type OfflineSaleInput struct {
	Items         []OfflineItemInput `json:"items" binding:"required,min=1"`
	OfflineNote   string             `json:"offline_note"`
	IsVIP         bool               `json:"is_vip"`
	PaymentMethod string             `json:"payment_method"` // "cash" or "card"
	CardType      string             `json:"card_type"`      // "humo", "uzcard", "visa", "mastercard"
}

func CreateOfflineSale(c *gin.Context) {
	var input OfflineSaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	// VIP sales are free, so payment method is irrelevant. Otherwise validate it.
	paymentMethod := input.PaymentMethod
	cardType := ""
	if input.IsVIP {
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

	order := models.Order{
		UserID:        nil,
		WorkerID:      workerIDPtr,
		Status:        "delivered",
		Phone:         "offline",
		OrderCode:     generateOrderCode(),
		IsOffline:     true,
		IsVIP:         input.IsVIP,
		PaymentMethod: paymentMethod,
		CardType:      cardType,
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

		// Products are sold only by full pack/capsule.
		var price float64
		if !input.IsVIP {
			price = product.PricePerPack * float64(item.Quantity)
		}

		orderItem := models.OrderItem{
			OrderID:          order.ID,
			ProductID:        item.ProductID,
			Quantity:         item.Quantity,
			OriginalQuantity: item.Quantity,
			UnitType:         "pack",
			Price:            price,
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
