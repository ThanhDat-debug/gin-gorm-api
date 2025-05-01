package controllers

import (
	"net/http"

	"github.com/ThanhDat-debug/gin-gorm-api/config"
	"github.com/ThanhDat-debug/gin-gorm-api/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Posts").Find(&users)
	c.JSON(http.StatusOK, users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}
