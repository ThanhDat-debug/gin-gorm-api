package routes

import (
	"github.com/ThanhDat-debug/gin-gorm-api/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/users", controllers.CreateUser)

		v1.GET("/posts", controllers.GetPosts)
		v1.POST("/posts", controllers.CreatePost)
	}
}
