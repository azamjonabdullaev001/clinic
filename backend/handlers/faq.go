package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FAQInput struct {
	Question string   `json:"question" binding:"required"`
	Answers  []string `json:"answers" binding:"required"`
}

func normalizeAnswers(raw []string) []models.FAQAnswer {
	answers := make([]models.FAQAnswer, 0, len(raw))
	position := 0
	for _, item := range raw {
		text := strings.TrimSpace(item)
		if text == "" {
			continue
		}
		answers = append(answers, models.FAQAnswer{Text: text, Position: position})
		position++
	}
	return answers
}

func GetFAQs(c *gin.Context) {
	var faqs []models.FAQ
	err := database.DB.
		Preload("Answers", func(db *gorm.DB) *gorm.DB {
			return db.Order("position asc, id asc")
		}).
		Order("created_at desc").
		Find(&faqs).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при загрузке FAQ"})
		return
	}

	c.JSON(http.StatusOK, faqs)
}

func CreateFAQ(c *gin.Context) {
	var input FAQInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите вопрос и хотя бы один ответ"})
		return
	}

	question := strings.TrimSpace(input.Question)
	answers := normalizeAnswers(input.Answers)
	if question == "" || len(answers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите вопрос и хотя бы один ответ"})
		return
	}

	faq := models.FAQ{Question: question, Answers: answers}
	if err := database.DB.Create(&faq).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании FAQ"})
		return
	}

	database.DB.Preload("Answers", func(db *gorm.DB) *gorm.DB {
		return db.Order("position asc, id asc")
	}).First(&faq, faq.ID)
	c.JSON(http.StatusCreated, faq)
}

func UpdateFAQ(c *gin.Context) {
	id := c.Param("id")
	var faq models.FAQ
	if err := database.DB.First(&faq, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "FAQ не найден"})
		return
	}

	var input FAQInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите вопрос и хотя бы один ответ"})
		return
	}

	question := strings.TrimSpace(input.Question)
	answers := normalizeAnswers(input.Answers)
	if question == "" || len(answers) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Введите вопрос и хотя бы один ответ"})
		return
	}

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		faq.Question = question
		if err := tx.Save(&faq).Error; err != nil {
			return err
		}
		if err := tx.Where("faq_id = ?", faq.ID).Delete(&models.FAQAnswer{}).Error; err != nil {
			return err
		}
		for _, answer := range answers {
			answer.FAQID = faq.ID
			if err := tx.Create(&answer).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении FAQ"})
		return
	}

	database.DB.Preload("Answers", func(db *gorm.DB) *gorm.DB {
		return db.Order("position asc, id asc")
	}).First(&faq, faq.ID)
	c.JSON(http.StatusOK, faq)
}

func DeleteFAQ(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.FAQ{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении FAQ"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "FAQ удален"})
}
