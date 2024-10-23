package routes

import (
    "github.com/Makeyabe/Home_Backend/controllers"
    "github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.RouterGroup) {
    imageRoutes := router.Group("/images")
    {
        imageRoutes.POST("/", controllers.CreateImage)
        imageRoutes.GET("/", controllers.GetImages)
        imageRoutes.GET("/:id", controllers.GetImage)
        imageRoutes.PUT("/:id", controllers.UpdateImage)
        imageRoutes.DELETE("/:id", controllers.DeleteImage)
    }
}
