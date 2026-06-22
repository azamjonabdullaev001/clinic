package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ProductInput struct {
	Name            string  `json:"name" binding:"required"`
	Description     string  `json:"description"`
	Category        string  `json:"category"`
	// QuantityPerPack is OPTIONAL: 0 means the product has no флакон and is sold only by
	// the piece (ointments / mixtures / liquids), priced at price_per_pill per unit.
	QuantityPerPack int     `json:"quantity_per_pack" binding:"omitempty,min=1"`
	PricePerPill    float64 `json:"price_per_pill" binding:"required,min=0"`
	StockQuantity   int     `json:"stock_quantity"`
	// OnlineAvailable is a pointer so an omitted field keeps the existing value (on update)
	// or defaults to true (on create) instead of silently hiding the product.
	OnlineAvailable *bool `json:"online_available"`
}

// GetProducts returns EVERY product. Used by admin/worker panels where offline sales must
// always see the full catalogue, regardless of the online-availability flag.
func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Order("created_at desc").Find(&products)

	for i := range products {
		products[i].ComputePackPrice()
	}

	c.JSON(http.StatusOK, products)
}

// GetPublicProducts returns only products flagged as available online. This backs the public
// landing page so unchecked products are neither shown nor orderable online.
func GetPublicProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Where("online_available = ?", true).Order("created_at desc").Find(&products)

	for i := range products {
		products[i].ComputePackPrice()
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Препарат не найден"})
		return
	}
	product.ComputePackPrice()
	c.JSON(http.StatusOK, product)
}

func CreateProduct(c *gin.Context) {
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заполните все обязательные поля"})
		return
	}

	product := models.Product{
		Name:            input.Name,
		Description:     input.Description,
		Category:        input.Category,
		QuantityPerPack: input.QuantityPerPack,
		PricePerPill:    input.PricePerPill,
		StockQuantity:   input.StockQuantity,
		OnlineAvailable: true, // new products are visible online by default
	}
	if input.OnlineAvailable != nil {
		product.OnlineAvailable = *input.OnlineAvailable
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании препарата"})
		return
	}

	product.ComputePackPrice()
	BroadcastProducts()
	c.JSON(http.StatusCreated, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Препарат не найден"})
		return
	}

	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заполните все обязательные поля"})
		return
	}

	product.Name = input.Name
	product.Description = input.Description
	product.Category = input.Category
	product.QuantityPerPack = input.QuantityPerPack
	product.PricePerPill = input.PricePerPill
	product.StockQuantity = input.StockQuantity
	if input.OnlineAvailable != nil {
		product.OnlineAvailable = *input.OnlineAvailable
	}

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	BroadcastStock(product.ID, product.StockQuantity)
	BroadcastProducts()
	product.ComputePackPrice()
	c.JSON(http.StatusOK, product)
}

type OnlineAvailabilityInput struct {
	OnlineAvailable bool `json:"online_available"`
}

// SetProductOnlineAvailability flips the single online-availability flag for a product.
// This is the checkbox in the warehouse list — it changes nothing about offline sales.
func SetProductOnlineAvailability(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Препарат не найден"})
		return
	}

	var input OnlineAvailabilityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверные данные"})
		return
	}

	if err := database.DB.Model(&product).UpdateColumn("online_available", input.OnlineAvailable).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	product.OnlineAvailable = input.OnlineAvailable
	product.ComputePackPrice()
	BroadcastProducts()
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Препарат не найден"})
		return
	}

	// Delete related order items first to avoid foreign key constraint
	database.DB.Where("product_id = ?", id).Delete(&models.OrderItem{})

	if err := database.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}
	BroadcastProducts()
	c.JSON(http.StatusOK, gin.H{"message": "Препарат удален"})
}

func UploadProductImage(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Препарат не найден"})
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Изображение не загружено"})
		return
	}

	if file.Size > 10<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл слишком большой (макс. 10MB)"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowed[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неподдерживаемый формат изображения"})
		return
	}

	filename := fmt.Sprintf("%s%s", uuid.New().String(), ext)
	savePath := filepath.Join("uploads", filename)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return
	}

	product.ImagePath = "/uploads/" + filename
	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	product.ComputePackPrice()
	BroadcastProducts()
	c.JSON(http.StatusOK, product)
}
