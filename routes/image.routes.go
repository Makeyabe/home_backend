package routes

import (
    "github.com/Makeyabe/Home_Backend/controllers"
    "github.com/gin-gonic/gin"
)

func SetupImageRoutes(router *gin.RouterGroup, ImageController *controllers.ImageController) {
    imageRoutes := router.Group("/images")
    {
        imageRoutes.POST("/", ImageController.CreateImage)
        // imageRoutes.GET("/", ImageController.GetImages)
        // imageRoutes.GET("/:id", ImageController.GetImage)
        // imageRoutes.PUT("/:id", ImageController.UpdateImage)
        // imageRoutes.DELETE("/:id", ImageController.DeleteImage)
    }
}
