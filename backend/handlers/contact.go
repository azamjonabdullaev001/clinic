package handlers

import (
	"bytes"
	"clinic-backend/config"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

type ContactInput struct {
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Message string `json:"message"`
}

var contactPhoneRegex = regexp.MustCompile(`^\+?[0-9\s()\-]{7,20}$`)
const contactTelegramChatID = "1941772742"

func SendContactMessage(c *gin.Context) {
	var input ContactInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заполните имя и телефон"})
		return
	}

	name := strings.TrimSpace(input.Name)
	phone := strings.TrimSpace(input.Phone)
	message := strings.TrimSpace(input.Message)

	if name == "" || phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Заполните имя и телефон"})
		return
	}

	if !contactPhoneRegex.MatchString(phone) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат номера телефона"})
		return
	}

	cfg := config.Load()
	if cfg.TelegramBotToken == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "Telegram бот не настроен"})
		return
	}

	var sb strings.Builder
	sb.WriteString("❓ ВОПРОС ОТ ПОЛЬЗОВАТЕЛЯ\n\n")
	sb.WriteString(fmt.Sprintf("👤 Имя: %s\n", name))
	sb.WriteString(fmt.Sprintf("📱 Телефон: %s\n", phone))
	if message != "" {
		sb.WriteString(fmt.Sprintf("💬 Сообщение: %s\n", message))
	} else {
		sb.WriteString("💬 Сообщение: (не указано)\n")
	}

	payload := map[string]string{
		"chat_id": contactTelegramChatID,
		"text":    sb.String(),
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Contact telegram marshal error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отправить сообщение"})
		return
	}

	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", cfg.TelegramBotToken)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("Contact telegram send error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отправить сообщение"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode >= http.StatusBadRequest {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отправить сообщение"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
