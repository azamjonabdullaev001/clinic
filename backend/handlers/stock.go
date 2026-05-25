package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetWorkerStock returns the logged-in worker's personal inventory.
func GetWorkerStock(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	var stock []models.WorkerStock
	database.DB.Where("worker_id = ?", workerID).
		Preload("Product").
		Order("updated_at desc").
		Find(&stock)
	for i := range stock {
		stock[i].Product.ComputePackPrice()
	}
	c.JSON(http.StatusOK, stock)
}

type AddStockInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

// AddWorkerStock adds incoming goods to the worker's own inventory.
func AddWorkerStock(c *gin.Context) {
	workerID, _ := c.Get("workerID")
	wid := workerID.(uint)

	var input AddStockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	var product models.Product
	if database.DB.First(&product, input.ProductID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Товар не найден"})
		return
	}

	var ws models.WorkerStock
	err := database.DB.Where("worker_id = ? AND product_id = ?", wid, input.ProductID).First(&ws).Error
	var saveErr error
	if err == gorm.ErrRecordNotFound {
		ws = models.WorkerStock{WorkerID: wid, ProductID: input.ProductID, Quantity: input.Quantity}
		if ws.Quantity < 0 {
			ws.Quantity = 0
		}
		saveErr = database.DB.Create(&ws).Error
	} else {
		ws.Quantity += input.Quantity
		if ws.Quantity < 0 {
			ws.Quantity = 0
		}
		saveErr = database.DB.Save(&ws).Error
	}
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении склада"})
		return
	}

	database.DB.Preload("Product").First(&ws, ws.ID)
	ws.Product.ComputePackPrice()
	c.JSON(http.StatusOK, ws)
}

// reserveWorkerStock verifies a worker has enough stock for the given items and
// decrements it. Quantities are aggregated per product so the same product added
// as several lines is checked correctly. Returns an error naming the first
// out-of-stock product.
func reserveWorkerStock(tx *gorm.DB, workerID uint, items []OfflineItemInput) error {
	need := map[uint]int{}
	for _, item := range items {
		need[item.ProductID] += item.Quantity
	}
	for productID, qty := range need {
		var ws models.WorkerStock
		err := tx.Where("worker_id = ? AND product_id = ?", workerID, productID).First(&ws).Error
		if err != nil || ws.Quantity < qty {
			var p models.Product
			tx.First(&p, productID)
			name := p.Name
			if name == "" {
				name = fmt.Sprintf("#%d", productID)
			}
			return fmt.Errorf("Недостаточно на складе: %s", name)
		}
	}
	for productID, qty := range need {
		tx.Model(&models.WorkerStock{}).
			Where("worker_id = ? AND product_id = ?", workerID, productID).
			UpdateColumn("quantity", gorm.Expr("quantity - ?", qty))
	}
	return nil
}

// restockWorker adds quantity back to a worker's stock (used for returns).
func restockWorker(workerID, productID uint, quantity int) {
	if quantity <= 0 {
		return
	}
	var ws models.WorkerStock
	err := database.DB.Where("worker_id = ? AND product_id = ?", workerID, productID).First(&ws).Error
	if err == gorm.ErrRecordNotFound {
		database.DB.Create(&models.WorkerStock{WorkerID: workerID, ProductID: productID, Quantity: quantity})
	} else {
		database.DB.Model(&models.WorkerStock{}).Where("id = ?", ws.ID).
			UpdateColumn("quantity", gorm.Expr("quantity + ?", quantity))
	}
}

// decrementWorkerStockNoBlock reduces a worker's stock for delivered order items
// without blocking (quantity floors at 0). Used when a worker hands over an
// online/doctor order — the goods leave that worker's warehouse.
func decrementWorkerStockNoBlock(workerID uint, items []models.OrderItem) {
	for _, it := range items {
		if it.Quantity <= 0 {
			continue
		}
		var ws models.WorkerStock
		err := database.DB.Where("worker_id = ? AND product_id = ?", workerID, it.ProductID).First(&ws).Error
		if err == gorm.ErrRecordNotFound {
			database.DB.Create(&models.WorkerStock{WorkerID: workerID, ProductID: it.ProductID, Quantity: 0})
			continue
		}
		newQty := ws.Quantity - it.Quantity
		if newQty < 0 {
			newQty = 0
		}
		database.DB.Model(&models.WorkerStock{}).Where("id = ?", ws.ID).UpdateColumn("quantity", newQty)
	}
}
