package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetMarketologDebt returns a marketolog's ALL-TIME debt (the retail value of everything
// sold under their name), how much they've already paid back, the remaining balance, and a
// per-product breakdown. Used by the cashier to settle marketolog debt.
func GetMarketologDebt(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var worker models.Worker
	if database.DB.First(&worker, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Маркетолог не найден"})
		return
	}

	// All-time debt aggregation. This MUST cover every order ever made under the
	// marketolog (no LIMIT), so instead of loading them into Go memory — which OOMs at
	// scale — we aggregate in SQL with a single GROUP BY that stays constant-memory.
	type prodAgg struct {
		ProductName string
		Revenue     float64
		Pieces      int
		Capsules    int
	}
	var rows []prodAgg
	database.DB.
		Table("order_items AS oi").
		Select(`p.name AS product_name,
			COALESCE(SUM(oi.price), 0) AS revenue,
			COALESCE(SUM(CASE WHEN oi.unit_type = 'piece' THEN oi.quantity ELSE 0 END), 0) AS pieces,
			COALESCE(SUM(CASE WHEN oi.unit_type <> 'piece' THEN oi.quantity ELSE 0 END), 0) AS capsules`).
		Joins("JOIN orders o ON o.id = oi.order_id").
		Joins("LEFT JOIN products p ON p.id = oi.product_id").
		Where("o.marketolog_id = ? AND o.status <> ? AND o.is_deleted = ? AND oi.quantity > 0", id, "cancelled", false).
		Group("p.name").
		Scan(&rows)

	var debt float64
	capsules, pieces := 0, 0
	prods := make([]MarketologProduct, 0, len(rows))
	for _, r := range rows {
		debt += r.Revenue
		capsules += r.Capsules
		pieces += r.Pieces
		prods = append(prods, MarketologProduct{
			ProductName: r.ProductName,
			Revenue:     r.Revenue,
			Pieces:      r.Pieces,
			Capsules:    r.Capsules,
		})
	}

	var totalOrders int64
	database.DB.Model(&models.Order{}).
		Where("marketolog_id = ? AND status <> ? AND is_deleted = ?", id, "cancelled", false).
		Count(&totalOrders)

	var paid float64
	database.DB.Model(&models.MarketologPayment{}).Where("marketolog_id = ?", id).
		Select("COALESCE(SUM(amount), 0)").Scan(&paid)

	var payments []models.MarketologPayment
	database.DB.Where("marketolog_id = ?", id).Order("created_at desc").Limit(100).Find(&payments)

	c.JSON(http.StatusOK, gin.H{
		"marketolog":     gin.H{"id": worker.ID, "name": worker.Name},
		"debt":           debt,
		"paid":           paid,
		"remaining":      debt - paid,
		"total_orders":   totalOrders,
		"total_capsules": capsules,
		"total_pieces":   pieces,
		"products":       prods,
		"payments":       payments,
	})
}

// AddMarketologPayment records a debt repayment from a marketolog (always a transfer/ХР).
func AddMarketologPayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var input struct {
		Amount float64 `json:"amount"`
		Note   string  `json:"note"`
	}
	if c.ShouldBindJSON(&input) != nil || input.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Укажите сумму оплаты"})
		return
	}
	// Verify target is actually a manager/marketolog worker
	var target models.Worker
	if err := database.DB.First(&target, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Маркетолог не найден"})
		return
	}
	if target.Role != "manager" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Выплаты доступны только для менеджеров"})
		return
	}

	var workerIDPtr *uint
	if wid, ok := c.Get("workerID"); ok {
		if w, ok := wid.(uint); ok {
			workerIDPtr = &w
		}
	}
	payment := models.MarketologPayment{
		MarketologID: uint(id),
		Amount:       input.Amount,
		WorkerID:     workerIDPtr,
		Note:         input.Note,
	}
	if database.DB.Create(&payment).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении"})
		return
	}
	c.JSON(http.StatusCreated, payment)
}
