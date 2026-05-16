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
)

func GetNewsPosts(c *gin.Context) {
	var posts []models.NewsPost
	database.DB.Order("created_at desc").Find(&posts)
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
	c.JSON(http.StatusOK, post)
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
