package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// parseHistoryRange reads ?from=YYYY-MM-DD&to=YYYY-MM-DD and returns a [start, end) range in
// local time. `to` is inclusive (end = to + 24h). Missing bounds fall back to a wide window so
// the caller still gets the full history when no date is chosen.
func parseHistoryRange(c *gin.Context) (time.Time, time.Time) {
	loc := time.Local
	now := time.Now().In(loc)
	start := time.Date(2000, 1, 1, 0, 0, 0, 0, loc)
	end := now.Add(24 * time.Hour)
	if f := c.Query("from"); f != "" {
		if p, err := time.ParseInLocation("2006-01-02", f, loc); err == nil {
			start = p
		}
	}
	if t := c.Query("to"); t != "" {
		if p, err := time.ParseInLocation("2006-01-02", t, loc); err == nil {
			end = p.Add(24 * time.Hour)
		}
	}
	return start, end
}

// sumOrderRevenue adds up the retail value of an order's non-removed items.
func sumOrderRevenue(orders []models.Order) float64 {
	var revenue float64
	for _, o := range orders {
		for _, it := range o.Items {
			if it.Quantity > 0 {
				revenue += it.Price
			}
		}
	}
	return revenue
}

// GetDoctorHistoryPickup returns ONE doctor's order history within a date range, for the
// pickup cashier's per-person history view. Orders are matched by doctor_id OR by the doctor's
// name appearing in referred_by — the same rule GetDoctorStats uses.
func GetDoctorHistoryPickup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var doctor models.Doctor
	if database.DB.First(&doctor, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Доктор не найден"})
		return
	}
	start, end := parseHistoryRange(c)

	var orders []models.Order
	database.DB.
		Where("(doctor_id = ? OR referred_by LIKE ?) AND status <> ? AND is_deleted = ? AND created_at >= ? AND created_at < ?",
			doctor.ID, "%"+doctor.Name+"%", "cancelled", false, start, end).
		Preload("Items.Product").
		Order("created_at desc").
		Limit(500).
		Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"person":       gin.H{"id": doctor.ID, "name": doctor.Name, "type": "doctor"},
		"total_orders": len(orders),
		"revenue":      sumOrderRevenue(orders),
		"orders":       orders,
	})
}

// GetMarketologHistoryPickup returns ONE marketolog's (manager worker) order history within a
// date range, for the pickup cashier's per-person history view.
func GetMarketologHistoryPickup(c *gin.Context) {
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
	start, end := parseHistoryRange(c)

	var orders []models.Order
	database.DB.
		Where("marketolog_id = ? AND status <> ? AND is_deleted = ? AND created_at >= ? AND created_at < ?",
			worker.ID, "cancelled", false, start, end).
		Preload("Items.Product").
		Order("created_at desc").
		Limit(500).
		Find(&orders)

	c.JSON(http.StatusOK, gin.H{
		"person":       gin.H{"id": worker.ID, "name": worker.Name, "type": "marketolog"},
		"total_orders": len(orders),
		"revenue":      sumOrderRevenue(orders),
		"orders":       orders,
	})
}
