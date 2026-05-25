package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StockRow struct {
	ProductID uint           `json:"product_id"`
	Product   models.Product `json:"product"`
	Quantity  int            `json:"quantity"`
}

// GetGlobalStock returns every product with its current warehouse quantity.
func GetGlobalStock(c *gin.Context) {
	var products []models.Product
	database.DB.Order("name asc").Find(&products)
	rows := make([]StockRow, len(products))
	for i := range products {
		products[i].ComputePackPrice()
		rows[i] = StockRow{ProductID: products[i].ID, Product: products[i], Quantity: products[i].StockQuantity}
	}
	c.JSON(http.StatusOK, rows)
}

type AddStockInput struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

// AddProductStock increases the shared warehouse quantity of a product.
func AddProductStock(c *gin.Context) {
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
	newQty := product.StockQuantity + input.Quantity
	if newQty < 0 {
		newQty = 0
	}
	if err := database.DB.Model(&models.Product{}).Where("id = ?", product.ID).
		Update("stock_quantity", newQty).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении склада"})
		return
	}
	BroadcastStock(product.ID, newQty)
	c.JSON(http.StatusOK, gin.H{"product_id": product.ID, "quantity": newQty})
}

// reserveProductStock checks the shared warehouse has enough of each item and
// decrements it. Quantities are aggregated per product. Blocks if insufficient.
func reserveProductStock(tx *gorm.DB, items []OfflineItemInput) error {
	need := map[uint]int{}
	for _, item := range items {
		need[item.ProductID] += item.Quantity
	}
	for pid, qty := range need {
		var p models.Product
		if tx.First(&p, pid).Error != nil {
			return fmt.Errorf("Товар не найден")
		}
		if p.StockQuantity < qty {
			return fmt.Errorf("Недостаточно на складе: %s (осталось %d)", p.Name, p.StockQuantity)
		}
	}
	for pid, qty := range need {
		tx.Model(&models.Product{}).Where("id = ?", pid).
			UpdateColumn("stock_quantity", gorm.Expr("stock_quantity - ?", qty))
	}
	return nil
}

// checkProductStock verifies the shared warehouse has enough for the items
// without decrementing (used at online order placement).
func checkProductStock(tx *gorm.DB, items []OrderItemInput) error {
	need := map[uint]int{}
	for _, item := range items {
		need[item.ProductID] += item.Quantity
	}
	for pid, qty := range need {
		var p models.Product
		if tx.First(&p, pid).Error != nil {
			return fmt.Errorf("Препарат не найден")
		}
		if p.StockQuantity < qty {
			return fmt.Errorf("Недостаточно на складе: %s (осталось %d)", p.Name, p.StockQuantity)
		}
	}
	return nil
}

// restockProduct returns quantity to the shared warehouse (used for returns).
func restockProduct(productID uint, qty int) {
	if qty <= 0 {
		return
	}
	database.DB.Model(&models.Product{}).Where("id = ?", productID).
		UpdateColumn("stock_quantity", gorm.Expr("stock_quantity + ?", qty))
	broadcastProductStock(productID)
}
