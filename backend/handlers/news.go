package handlers

import (
	"clinic-backend/database"
	"clinic-backend/models"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetNewsPosts(c *gin.Context) {
	var posts []models.NewsPost
	database.DB.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("position asc")
	}).Order("created_at desc").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func CreateNewsPost(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	videoURL := c.PostForm("video_url")

	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	post := models.NewsPost{
		Title:       title,
		Description: description,
		VideoURL:    videoURL,
	}

	file, err := c.FormFile("image")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("news_%d%s", time.Now().UnixNano(), ext)
		savePath := filepath.Join("uploads", filename)
		if err := os.MkdirAll("uploads", 0755); err == nil {
			if err := c.SaveUploadedFile(file, savePath); err == nil {
				post.ImagePath = "/uploads/" + filename
			}
		}
	}

	if err := database.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create news post"})
		return
	}

	// Handle multiple additional images
	form, _ := c.MultipartForm()
	if form != nil {
		files := form.File["images"]
		for i, imgFile := range files {
			ext := filepath.Ext(imgFile.Filename)
			filename := fmt.Sprintf("news_%d_%d%s", time.Now().UnixNano(), i, ext)
			savePath := filepath.Join("uploads", filename)
			if err := c.SaveUploadedFile(imgFile, savePath); err == nil {
				database.DB.Create(&models.NewsImage{
					NewsPostID: post.ID,
					ImagePath:  "/uploads/" + filename,
					Position:   i,
				})
			}
		}
	}

	database.DB.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("position asc")
	}).First(&post, post.ID)
	c.JSON(http.StatusCreated, post)
}

func UpdateNewsPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var post models.NewsPost
	if err := database.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	if title := c.PostForm("title"); title != "" {
		post.Title = title
	}
	post.Description = c.PostForm("description")
	post.VideoURL = c.PostForm("video_url")

	file, err := c.FormFile("image")
	if err == nil {
		ext := filepath.Ext(file.Filename)
		filename := fmt.Sprintf("news_%d%s", time.Now().UnixNano(), ext)
		savePath := filepath.Join("uploads", filename)
		if err := c.SaveUploadedFile(file, savePath); err == nil {
			post.ImagePath = "/uploads/" + filename
		}
	}

	database.DB.Save(&post)

	// Handle multiple additional images
	form, _ := c.MultipartForm()
	if form != nil {
		files := form.File["images"]
		for i, imgFile := range files {
			ext := filepath.Ext(imgFile.Filename)
			filename := fmt.Sprintf("news_%d_%d%s", time.Now().UnixNano(), i, ext)
			savePath := filepath.Join("uploads", filename)
			if err := c.SaveUploadedFile(imgFile, savePath); err == nil {
				database.DB.Create(&models.NewsImage{
					NewsPostID: post.ID,
					ImagePath:  "/uploads/" + filename,
					Position:   i,
				})
			}
		}
	}

	database.DB.Preload("Images", func(db *gorm.DB) *gorm.DB {
		return db.Order("position asc")
	}).First(&post, post.ID)
	c.JSON(http.StatusOK, post)
}

func DeleteNewsImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var img models.NewsImage
	if err := database.DB.First(&img, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	// Remove file from disk (ignore errors)
	if img.ImagePath != "" {
		os.Remove("." + img.ImagePath)
	}
	database.DB.Delete(&img)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func DeleteNewsPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	database.DB.Delete(&models.NewsPost{}, id)
	c.JSON(http.StatusOK, gin.H{"ok": true})
}
