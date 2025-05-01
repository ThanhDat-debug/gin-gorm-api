package main

import (
	"github.com/ThanhDat-debug/gin-gorm-api/config"
	"github.com/ThanhDat-debug/gin-gorm-api/models"
	"github.com/ThanhDat-debug/gin-gorm-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectDatabase()

	config.DB.AutoMigrate(&models.User{}, &models.Post{})

	routes.SetupRoutes(router)

	router.Run(":8080")
}
