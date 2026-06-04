package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"encoding/json"
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

// PaymentSplit is one method's share of a sale paid via several methods at once.
type PaymentSplit struct {
	Method string  `json:"method"` // cash, terminal, cassa1, click, transfer
	Amount float64 `json:"amount"`
}

type OfflineSaleInput struct {
	Items         []OfflineItemInput `json:"items" binding:"required,min=1"`
	OfflineNote     string             `json:"offline_note"`
	IsVIP           bool               `json:"is_vip"`
	PaymentMethod   string             `json:"payment_method"`   // "cash", "terminal", "card"
	CardType        string             `json:"card_type"`        // "cassa1", "click", "transfer"
	PaymentSplits   []PaymentSplit     `json:"payment_splits"`   // when paid via several methods at once
	DiscountPercent float64            `json:"discount_percent"` // 0..100 — cashier-typed discount on the whole sale
	ReferredBy    string             `json:"referred_by"`    // doctor who referred the patient
	SalesChannel  string             `json:"sales_channel"`  // marketplace for manager sales
	MarketologID  *uint              `json:"marketolog_id"`  // marketolog the debt-sale is assigned to
}

// canonicalPayment maps a split method to the legacy payment_method/card_type pair so
// single-method sales keep behaving exactly as before.
func canonicalPayment(method string) (string, string) {
	switch method {
	case "cassa1", "click", "transfer":
		return "card", method
	default: // cash, terminal
		return method, ""
	}
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

	// Resolve the marketolog this sale belongs to (debt). The cashier passes the
	// chosen marketolog id; a manager selling in their own panel is the marketolog.
	var marketologIDPtr *uint
	if input.MarketologID != nil && *input.MarketologID > 0 {
		marketologIDPtr = input.MarketologID
	} else if input.SalesChannel != "" && workerIDPtr != nil {
		marketologIDPtr = workerIDPtr
	}
	isDebt := marketologIDPtr != nil

	// Determine payment method. Marketolog/debt sales are settled later; "own patient"
	// sales are free; otherwise the cashier picks cash/terminal/card.
	paymentMethod := input.PaymentMethod
	cardType := ""
	paymentSplits := ""
	if isDebt {
		paymentMethod = "marketplace"
	} else if input.IsVIP {
		paymentMethod = ""
	} else {
		// Collect the non-zero split lines (an order may be paid via several methods).
		var nonZero []PaymentSplit
		for _, s := range input.PaymentSplits {
			if s.Amount > 0 && s.Method != "" {
				nonZero = append(nonZero, s)
			}
		}
		if len(nonZero) > 0 {
			if len(nonZero) == 1 {
				paymentMethod, cardType = canonicalPayment(nonZero[0].Method)
			} else {
				paymentMethod, cardType = "mixed", ""
			}
			if b, err := json.Marshal(nonZero); err == nil {
				paymentSplits = string(b)
			}
		} else {
			if paymentMethod == "" {
				paymentMethod = "cash"
			}
			if !validOfflinePayment(paymentMethod) {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный способ оплаты"})
				return
			}
			if paymentMethod == "card" {
				cardType = input.CardType
			}
		}
	}

	// Clamp the cashier-typed discount to a sane range. VIP and marketolog sales never
	// get a discount applied (VIP is already free, marketolog is debt).
	discount := input.DiscountPercent
	if isDebt || input.IsVIP {
		discount = 0
	}
	if discount < 0 {
		discount = 0
	}
	if discount > 100 {
		discount = 100
	}
	discountFactor := 1 - discount/100

	// Offline & marketolog sales draw from the shared warehouse and cannot exceed it.
	if err := reserveProductStock(tx, input.Items); err != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{
		UserID:        nil,
		WorkerID:      workerIDPtr,
		MarketologID:  marketologIDPtr,
		Status:        "delivered",
		Phone:         "offline",
		OrderCode:     generateNurseCode(), // offline orders use a 5-digit code
		IsOffline:     true,
		IsVIP:         input.IsVIP,
		PaymentMethod:   paymentMethod,
		CardType:        cardType,
		PaymentSplits:   paymentSplits,
		DiscountPercent: discount,
		SalesChannel:    input.SalesChannel,
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
			// Cashier-typed discount reduces each item proportionally so analytics
			// naturally reflects the actual money received.
			price *= discountFactor
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
	BroadcastOrders()

	database.DB.Preload("Items.Product").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusCreated, order)
}
