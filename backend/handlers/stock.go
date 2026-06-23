package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"fmt"
	"net/http"
	"strconv"

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
	ProductID uint   `json:"product_id" binding:"required"`
	Quantity  int    `json:"quantity" binding:"required,min=1"`
	UnitType  string `json:"unit_type"` // "pack" (capsules) or "piece"
}

// pieceCount converts a quantity in the given unit to pieces for a product.
func pieceCount(p models.Product, quantity int, unitType string) int {
	if unitType == "piece" {
		return quantity
	}
	qpp := p.QuantityPerPack
	if qpp < 1 {
		qpp = 1
	}
	return quantity * qpp
}

// AddProductStock increases the shared warehouse quantity of a product (in pieces).
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
	pieces := pieceCount(product, input.Quantity, input.UnitType)
	if err := database.DB.Model(&models.Product{}).Where("id = ?", product.ID).
		UpdateColumn("stock_quantity", gorm.Expr("stock_quantity + ?", pieces)).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении склада"})
		return
	}
	database.DB.First(&product, product.ID)
	BroadcastStock(product.ID, product.StockQuantity)
	c.JSON(http.StatusOK, gin.H{"product_id": product.ID, "quantity": product.StockQuantity})
}

type SetStockInput struct {
	Quantity int    `json:"quantity"`   // exact amount to set (>= 0)
	UnitType string `json:"unit_type"`  // "pack" (flacons/capsules) or "piece"
}

// SetProductStock OVERWRITES the shared warehouse quantity of a product to an exact value
// (stored in pieces). Unlike AddProductStock (which increments), this is used to correct or
// recount stock — e.g. after a physical inventory. The amount is given in packs or pieces.
func SetProductStock(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
		return
	}
	var input SetStockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}
	if input.Quantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Количество не может быть отрицательным"})
		return
	}
	var product models.Product
	if database.DB.First(&product, id).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Товар не найден"})
		return
	}
	pieces := pieceCount(product, input.Quantity, input.UnitType)
	if err := database.DB.Model(&models.Product{}).Where("id = ?", product.ID).
		UpdateColumn("stock_quantity", pieces).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении склада"})
		return
	}
	BroadcastStock(product.ID, pieces)
	c.JSON(http.StatusOK, gin.H{"product_id": product.ID, "quantity": pieces})
}

// reserveProductStock checks the shared warehouse has enough pieces for each item
// and decrements it. Quantities are aggregated per product. Blocks if insufficient.
func reserveProductStock(tx *gorm.DB, items []OfflineItemInput) error {
	need := map[uint]int{} // pieces needed per product
	for _, item := range items {
		var p models.Product
		if tx.First(&p, item.ProductID).Error != nil {
			return fmt.Errorf("Товар не найден")
		}
		need[item.ProductID] += pieceCount(p, item.Quantity, item.UnitType)
	}
	for pid, pieces := range need {
		var p models.Product
		tx.First(&p, pid)
		if p.StockQuantity < pieces {
			return fmt.Errorf("Недостаточно на складе: %s (осталось %d шт)", p.Name, p.StockQuantity)
		}
	}
	for pid, pieces := range need {
		tx.Model(&models.Product{}).Where("id = ?", pid).
			UpdateColumn("stock_quantity", gorm.Expr("stock_quantity - ?", pieces))
	}
	return nil
}

// checkProductStock verifies the warehouse has enough pieces for online items
// (always whole capsules) without decrementing.
func checkProductStock(tx *gorm.DB, items []OrderItemInput) error {
	need := map[uint]int{}
	for _, item := range items {
		var p models.Product
		if tx.First(&p, item.ProductID).Error != nil {
			return fmt.Errorf("Препарат не найден")
		}
		need[item.ProductID] += pieceCount(p, item.Quantity, item.UnitType)
	}
	for pid, pieces := range need {
		var p models.Product
		tx.First(&p, pid)
		if p.StockQuantity < pieces {
			return fmt.Errorf("Недостаточно на складе: %s (осталось %d шт)", p.Name, p.StockQuantity)
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

// decrementProductStock removes pieces from the warehouse (clamped at 0) and broadcasts
// the new level. Used when an edit ADDS or increases items on an already-delivered order.
func decrementProductStock(productID uint, qty int) {
	if qty <= 0 {
		return
	}
	database.DB.Model(&models.Product{}).Where("id = ?", productID).
		UpdateColumn("stock_quantity", gorm.Expr("GREATEST(stock_quantity - ?, 0)", qty))
	broadcastProductStock(productID)
}
