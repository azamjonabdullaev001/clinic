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
	QuantityPerPack int     `json:"quantity_per_pack" binding:"required,min=1"`
	PricePerPill    float64 `json:"price_per_pill" binding:"required,min=0"`
}

func GetProducts(c *gin.Context) {
	var products []models.Product
	database.DB.Order("created_at desc").Find(&products)

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
		QuantityPerPack: input.QuantityPerPack,
		PricePerPill:    input.PricePerPill,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании препарата"})
		return
	}

	product.ComputePackPrice()
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
	product.QuantityPerPack = input.QuantityPerPack
	product.PricePerPill = input.PricePerPill

	if err := database.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
		return
	}

	product.ComputePackPrice()
	c.JSON(http.StatusOK, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении"})
		return
	}
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
	c.JSON(http.StatusOK, product)
}
