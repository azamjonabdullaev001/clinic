package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetManagerOrders returns only the marketplace sales created by the logged-in manager.
func GetManagerOrders(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	var orders []models.Order
	database.DB.Where("worker_id = ?", workerID).
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
