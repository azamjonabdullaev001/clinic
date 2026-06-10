package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var imageExts = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}

func saveUploadedImage(c *gin.Context, field string) (string, bool) {
	file, err := c.FormFile(field)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Изображение не загружено"})
		return "", false
	}
	if file.Size > 10<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Файл слишком большой (макс. 10MB)"})
		return "", false
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !imageExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неподдерживаемый формат изображения"})
		return "", false
	}
	filename := uuid.New().String() + ext
	if err := c.SaveUploadedFile(file, filepath.Join("uploads", filename)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении файла"})
		return "", false
	}
	return "/uploads/" + filename, true
}

func getSetting(key string) string {
	var s models.Setting
	if database.DB.First(&s, "key = ?", key).Error != nil {
		return ""
	}
	return s.Value
}

func setSetting(key, value string) {
	database.DB.Save(&models.Setting{Key: key, Value: value})
}

// GetPaymentQR returns the URL of the pharmacy's payment QR image (public).
func GetPaymentQR(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"url": getSetting("payment_qr")})
}

// UploadPaymentQR sets the single static payment QR image (admin).
func UploadPaymentQR(c *gin.Context) {
	url, ok := saveUploadedImage(c, "image")
	if !ok {
		return
	}
	setSetting("payment_qr", url)
	c.JSON(http.StatusOK, gin.H{"url": url})
}
