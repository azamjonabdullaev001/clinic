package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SupportMessageInput struct {
	Message string `json:"message" binding:"required"`
}

func ensureSupportThread(userID uint) (models.SupportThread, error) {
	var thread models.SupportThread
	err := database.DB.Where("user_id = ?", userID).First(&thread).Error
	if err == nil {
		return thread, nil
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		return thread, err
	}

	thread = models.SupportThread{UserID: userID}
	if err := database.DB.Create(&thread).Error; err != nil {
		return thread, err
	}
	return thread, nil
}

func GetUserSupportThread(c *gin.Context) {
	userID, _ := c.Get("userID")
	thread, err := ensureSupportThread(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке поддержки"})
		return
	}

	database.DB.Model(&models.SupportMessage{}).
		Where("thread_id = ? AND sender_role = ?", thread.ID, "admin").
		Update("read_by_user", true)

	database.DB.Preload("Messages", func(db *gorm.DB) *gorm.DB {
		return db.Order("created_at asc")
	}).First(&thread, thread.ID)

	c.JSON(http.StatusOK, thread)
}

func SendUserSupportMessage(c *gin.Context) {
	userID, _ := c.Get("userID")
	thread, err := ensureSupportThread(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отправке сообщения"})
		return
	}

	var input SupportMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите сообщение"})
		return
	}

	message := strings.TrimSpace(input.Message)
	if message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите сообщение"})
		return
	}

	supportMessage := models.SupportMessage{
		ThreadID:    thread.ID,
		SenderRole:  "user",
		Message:     message,
		ReadByUser:  true,
		ReadByAdmin: false,
	}
	if err := database.DB.Create(&supportMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отправке сообщения"})
		return
	}

	database.DB.Model(&thread).Update("updated_at", supportMessage.CreatedAt)
	c.JSON(http.StatusCreated, supportMessage)
}

func GetSupportThreads(c *gin.Context) {
	var threads []models.SupportThread
	err := database.DB.
		Preload("User").
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at desc")
		}).
		Order("updated_at desc").
		Find(&threads).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке обращений"})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func GetSupportThreadByID(c *gin.Context) {
	id := c.Param("id")
	var thread models.SupportThread
	err := database.DB.
		Preload("User").
		Preload("Messages", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at asc")
		}).
		First(&thread, id).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Обращение не найдено"})
		return
	}

	database.DB.Model(&models.SupportMessage{}).
		Where("thread_id = ? AND sender_role = ?", thread.ID, "user").
		Update("read_by_admin", true)

	c.JSON(http.StatusOK, thread)
}

func ReplySupportThread(c *gin.Context) {
	id := c.Param("id")
	var thread models.SupportThread
	if err := database.DB.First(&thread, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Обращение не найдено"})
		return
	}

	var input SupportMessageInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите сообщение"})
		return
	}

	message := strings.TrimSpace(input.Message)
	if message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите сообщение"})
		return
	}

	supportMessage := models.SupportMessage{
		ThreadID:    thread.ID,
		SenderRole:  "admin",
		Message:     message,
		ReadByUser:  false,
		ReadByAdmin: true,
	}
	if err := database.DB.Create(&supportMessage).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при отправке сообщения"})
		return
	}

	database.DB.Model(&thread).Update("updated_at", supportMessage.CreatedAt)
	c.JSON(http.StatusCreated, supportMessage)
}
