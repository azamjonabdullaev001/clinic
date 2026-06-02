package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ReturnOrderFull fully returns a delivered order: all pieces go back to the
// warehouse, revenue is removed (order is cancelled) and the reason is recorded.
func ReturnOrderFull(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}
	if order.Status != "delivered" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Полный возврат доступен только для выданных заказов"})
		return
	}
	var input struct {
		ReturnReason string `json:"return_reason"`
	}
	c.ShouldBindJSON(&input)
	reason := strings.TrimSpace(input.ReturnReason)
	if reason == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Укажите причину возврата"})
		return
	}

	for _, it := range order.Items {
		if it.Quantity <= 0 {
			continue
		}
		var p models.Product
		if database.DB.First(&p, it.ProductID).Error == nil {
			restockProduct(it.ProductID, pieceCount(p, it.Quantity, it.UnitType))
		}
	}

	order.Status = "cancelled"
	order.IsReturned = true
	order.ReturnReason = reason
	if wid, ok := c.Get("workerID"); ok {
		w := wid.(uint)
		order.WorkerID = &w
	}
	database.DB.Save(&order)
	BroadcastOrders()

	database.DB.Preload("Items.Product").First(&order, order.ID)
	for i := range order.Items {
		order.Items[i].Product.ComputePackPrice()
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrderItems(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}

	if order.Status == "cancelled" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Нельзя изменить отменённый заказ"})
		return
	}

	// Editing a delivered order is a return: the customer brought goods back.
	isReturn := order.Status == "delivered"

	var input struct {
		Items        []OfflineItemInput `json:"items" binding:"required,min=1"`
		ReturnReason string             `json:"return_reason"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Список товаров обязателен"})
		return
	}

	if isReturn && strings.TrimSpace(input.ReturnReason) == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Укажите причину возврата"})
		return
	}

	productCache := map[uint]models.Product{}
	getProduct := func(tx *gorm.DB, id uint) (models.Product, bool) {
		if p, ok := productCache[id]; ok {
			return p, true
		}
		var p models.Product
		if tx.First(&p, id).Error != nil {
			return p, false
		}
		productCache[id] = p
		return p, true
	}

	// Track original prescription and previous pieces (for the diff and for restock).
	origByProduct := map[uint]int{}
	origUnitByProduct := map[uint]string{}
	prevPiecesByProduct := map[uint]int{}
	for _, it := range order.Items {
		if it.OriginalQuantity > origByProduct[it.ProductID] {
			origByProduct[it.ProductID] = it.OriginalQuantity
			origUnitByProduct[it.ProductID] = it.UnitType
		}
		if p, ok := getProduct(database.DB, it.ProductID); ok {
			prevPiecesByProduct[it.ProductID] += pieceCount(p, it.Quantity, it.UnitType)
		}
	}
	boughtProducts := map[uint]bool{}
	newPiecesByProduct := map[uint]int{}

	tx := database.DB.Begin()

	if err := tx.Where("order_id = ?", order.ID).Delete(&models.OrderItem{}).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	for _, item := range input.Items {
		product, ok := getProduct(tx, item.ProductID)
		if !ok {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": "Товар не найден"})
			return
		}
		product.ComputePackPrice()
		boughtProducts[item.ProductID] = true

		unitType := item.UnitType
		if unitType != "piece" {
			unitType = "pack"
		}
		newPiecesByProduct[item.ProductID] += pieceCount(product, item.Quantity, unitType)

		var price float64
		if !order.IsVIP {
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
			OriginalQuantity: origByProduct[item.ProductID], // 0 means newly added at the till
			UnitType:         unitType,
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
			unit := origUnitByProduct[productID]
			if unit != "piece" {
				unit = "pack"
			}
			orderItem := models.OrderItem{
				OrderID:          order.ID,
				ProductID:        productID,
				Quantity:         0,
				OriginalQuantity: origQty,
				UnitType:         unit,
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

	// Any item edit is recorded so the admin can review what the cashier changed (red border).
	order.IsEdited = true
	database.DB.Model(&models.Order{}).Where("id = ?", order.ID).Update("is_edited", true)

	// On a return, flag the order and put the brought-back pieces back into the warehouse.
	if isReturn {
		order.IsReturned = true
		order.ReturnReason = strings.TrimSpace(input.ReturnReason)
		database.DB.Model(&models.Order{}).Where("id = ?", order.ID).
			Updates(map[string]interface{}{"is_returned": true, "return_reason": order.ReturnReason})
		for productID, prevPieces := range prevPiecesByProduct {
			returned := prevPieces - newPiecesByProduct[productID]
			if returned > 0 {
				restockProduct(productID, returned)
			}
		}
	}

	BroadcastOrders()

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
		"archived = false AND is_deleted = false AND ("+
			"(is_offline = false AND is_nurse_order = false AND status NOT IN ('delivered','cancelled')) "+
			"OR (is_nurse_order = true AND status NOT IN ('delivered','cancelled')) "+
			"OR worker_id = ?)",
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
		CardType           string `json:"card_type"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный статус"})
		return
	}

	validStatuses := map[string]bool{
		"pending":    true,
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
			if input.PaymentMethod == "card" {
				order.CardType = input.CardType
			} else {
				order.CardType = ""
			}
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
		// Online & doctor orders draw down the shared warehouse on delivery.
		// (Direct offline sales already decremented it at creation.)
		decrementStockForOrder(order)
	}
	BroadcastOrders()

	c.JSON(http.StatusOK, order)
}

// DeletePickupOrder soft-deletes an order from the pickup point. It disappears from
// the cashier's lists and from analytics, but stays visible on the admin panel with a
// black border so the admin can ask the cashier why it was removed. If the order was
// already delivered, its pieces are returned to the warehouse so stock stays exact.
func DeletePickupOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := database.DB.Preload("Items").First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заказ не найден"})
		return
	}
	if order.IsDeleted {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заказ уже удалён"})
		return
	}

	var input struct {
		Reason string `json:"reason"`
	}
	c.ShouldBindJSON(&input)
	reason := strings.TrimSpace(input.Reason)
	if reason == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Укажите причину удаления"})
		return
	}

	// A delivered order had already drawn down stock — return those pieces so the
	// closed-system reconciliation (stock + sold) stays exact after the deletion.
	if order.Status == "delivered" {
		for _, it := range order.Items {
			if it.Quantity <= 0 {
				continue
			}
			var p models.Product
			if database.DB.First(&p, it.ProductID).Error == nil {
				restockProduct(it.ProductID, pieceCount(p, it.Quantity, it.UnitType))
			}
		}
	}

	order.IsDeleted = true
	order.DeletedReason = reason
	if workerID, ok := c.Get("workerID"); ok {
		wid := workerID.(uint)
		var worker models.Worker
		if err := database.DB.First(&worker, wid).Error; err == nil {
			order.DeletedByName = worker.Name
			order.DeletedByRole = worker.Role
		}
	}
	database.DB.Save(&order)
	BroadcastOrders()

	c.JSON(http.StatusOK, gin.H{"message": "Заказ удалён"})
}
