package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdateOrderItems(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	if order.Status == "delivered" || order.Status == "cancelled" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Нельзя изменить завершённый заказ"})
		return
	}

	var input struct {
		Items []OfflineItemInput `json:"items" binding:"required,min=1"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Список товаров обязателен"})
		return
	}

	// Preserve the originally ordered quantity per product (the doctor's prescription),
	// so the admin can later see what was bought / added / removed.
	origByProduct := map[uint]int{}
	for _, it := range order.Items {
		if it.OriginalQuantity > origByProduct[it.ProductID] {
			origByProduct[it.ProductID] = it.OriginalQuantity
		}
	}
	boughtProducts := map[uint]bool{}
	for _, it := range input.Items {
		boughtProducts[it.ProductID] = true
	}

	tx := database.DB.Begin()

	if err := tx.Where("order_id = ?", order.ID).Delete(&models.OrderItem{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	// Bought items (products sold only by full capsule/pack; VIP = free).
	for _, item := range input.Items {
		var product models.Product
		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Товар не найден"})
			return
		}
		product.ComputePackPrice()

		var price float64
		if !order.IsVIP {
			price = product.PricePerPack * float64(item.Quantity)
		}

		orderItem := models.OrderItem{
			OrderID:          order.ID,
			ProductID:        item.ProductID,
			Quantity:         item.Quantity,
			OriginalQuantity: origByProduct[item.ProductID], // 0 means newly added at the till
			UnitType:         "pack",
			Price:            price,
		}
		if err := tx.Create(&orderItem).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
			return
		}
	}

	// Prescribed items that were not bought: keep them as a zero-quantity record so the
	// admin sees them marked as "убрано". Price 0, so they don't affect totals/analytics.
	for productID, origQty := range origByProduct {
		if origQty > 0 && !boughtProducts[productID] {
			orderItem := models.OrderItem{
				OrderID:          order.ID,
				ProductID:        productID,
				Quantity:         0,
				OriginalQuantity: origQty,
				UnitType:         "pack",
				Price:            0,
			}
			if err := tx.Create(&orderItem).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
				return
			}
		}
	}

	tx.Commit()

	database.DB.Preload("Items.Product").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusOK, order)
}

func GetPickupOrders(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	var orders []models.Order
	// Incoming (not yet finalized) online customer orders and doctor/nurse pre-orders are
	// shared with every pickup worker, so anyone can serve a walk-in. Everything that is
	// already finalized or is a direct offline sale stays private to its own worker.
	database.DB.Where(
		"(is_offline = false AND is_nurse_order = false AND status NOT IN ('delivered','cancelled')) "+
			"OR (is_nurse_order = true AND status NOT IN ('delivered','cancelled')) "+
			"OR worker_id = ?",
		workerID,
	).Preload("Items.Product").Preload("User").
		Order("created_at desc").
		Find(&orders)

	for i := range orders {
		for j := range orders[i].Items {
			orders[i].Items[j].Product.ComputePackPrice()
		}
	}

	c.JSON(http.StatusOK, orders)
}

func GetOrderByCode(c *gin.Context) {
	code := c.Param("code")
	var order models.Order
	if err := database.DB.Preload("Items.Product").Preload("User").
		Where("order_code = ?", code).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	c.JSON(http.StatusOK, order)
}

func UpdatePickupOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	var input struct {
		Status             string `json:"status" binding:"required"`
		CancellationReason string `json:"cancellation_reason"`
		PaymentMethod      string `json:"payment_method"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный статус"})
		return
	}

	validStatuses := map[string]bool{
		"pending":    true,
		"confirmed":  true,
		"shipped":    true,
		"in_transit": true,
		"delivered":  true,
		"cancelled":  true,
	}
	if !validStatuses[input.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный статус"})
		return
	}

	if input.Status == "cancelled" {
		reason := strings.TrimSpace(input.CancellationReason)
		if reason == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Укажите причину отмены"})
			return
		}
		order.CancellationReason = reason
		if workerID, ok := c.Get("workerID"); ok {
			wid := workerID.(uint)
			order.WorkerID = &wid // keep cancelled order private to the worker who cancelled it
			var worker models.Worker
			if err := database.DB.First(&worker, wid).Error; err == nil {
				order.CancelledByName = worker.Name
				order.CancelledByRole = worker.Role
			}
		}
	}

	if input.Status == "delivered" {
		if workerID, ok := c.Get("workerID"); ok {
			wid := workerID.(uint)
			order.WorkerID = &wid
		}
		// Record how the customer paid. Offline orders (doctor/nurse pre-orders) require an
		// explicit method chosen by the cashier; online orders are always paid online by card.
		if order.IsVIP {
			order.PaymentMethod = ""
		} else if order.IsOffline {
			if !validOfflinePayment(input.PaymentMethod) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Выберите способ оплаты"})
				return
			}
			order.PaymentMethod = input.PaymentMethod
		} else {
			order.PaymentMethod = "online"
		}
	}

	order.Status = input.Status
	database.DB.Save(&order)

	database.DB.Preload("Items.Product").Preload("User").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}

	if input.Status == "delivered" {
		go sendTelegramNotification(order)
		decrementStockForOrder(order)
	}

	c.JSON(http.StatusOK, order)
}
