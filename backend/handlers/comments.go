package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetProductComments returns the public comments for a product.
func GetProductComments(c *gin.Context) {
	id := c.Param("id")
	var comments []models.ProductComment
	database.DB.Where("product_id = ?", id).Order("created_at desc").Find(&comments)
	c.JSON(http.StatusOK, comments)
}

type CommentInput struct {
	Text string `json:"text" binding:"required"`
}

// AddProductComment lets a logged-in user leave a public comment on a product.
// The author name is taken from their account; anonymous users cannot comment.
func AddProductComment(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if database.DB.First(&product, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}

	var input CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Текст комментария обязателен"})
		return
	}
	text := strings.TrimSpace(input.Text)
	if text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Текст комментария обязателен"})
		return
	}

	name := "Пользователь"
	if uid, ok := c.Get("userID"); ok {
		var user models.User
		if database.DB.First(&user, uid).Error == nil {
			full := strings.TrimSpace(user.FirstName + " " + user.LastName)
			if full != "" {
				name = full
			} else if user.Phone != "" {
				name = user.Phone
			}
		}
	}

	comment := models.ProductComment{ProductID: product.ID, AuthorName: name, Text: text}
	if err := database.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при сохранении"})
		return
	}
	c.JSON(http.StatusCreated, comment)
}

// DeleteProductCommentByID lets an admin remove a comment.
func DeleteProductCommentByID(c *gin.Context) {
	database.DB.Delete(&models.ProductComment{}, c.Param("commentId"))
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
