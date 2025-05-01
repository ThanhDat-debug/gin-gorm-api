package controllers

import (
	"net/http"

	"github.com/ThanhDat-debug/gin-gorm-api/config"
	"github.com/ThanhDat-debug/gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}
