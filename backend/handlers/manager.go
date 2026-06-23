package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetManagerOrders returns all sales assigned to the logged-in marketolog —
// both their own and the ones a cashier recorded on their behalf.
func GetManagerOrders(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	// Per-marketolog list — grows slowly, and ManagerPanel sums an all-time stat over it,
	// so keep the bound generous. Still capped so it can never load an unbounded table.
	limit, offset := paginate(c, 2000, 5000)
	var orders []models.Order
	database.DB.Where("marketolog_id = ? AND archived = ? AND is_deleted = ?", workerID, false, false).
		Preload("Items.Product").
		Order("created_at desc").
		Limit(limit).Offset(offset).
		Find(&orders)

	for i := range orders {
		for j := range orders[i].Items {
			orders[i].Items[j].Product.ComputePackPrice()
		}
	}

	c.JSON(http.StatusOK, orders)
}

// GetMarketologs returns the list of marketolog workers (role "manager") so the
// cashier can record a sale on behalf of a specific marketolog.
func GetMarketologs(c *gin.Context) {
	var workers []models.Worker
	database.DB.Where("role = ?", "manager").Order("name asc").Find(&workers)
	result := make([]gin.H, len(workers))
	for i, w := range workers {
		result[i] = gin.H{"id": w.ID, "name": w.Name}
	}
	c.JSON(http.StatusOK, result)
}
