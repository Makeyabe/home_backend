package routes

import (
	"github.com/Makeyabe/Home_Backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FormResponseRoutes(r *gin.Engine, db *gorm.DB) {
	formRoutes := r.Group("/form-responses")
	{
		formRoutes.POST("/", func(c *gin.Context) { controllers.CreateFormResponse(c, db) })
		formRoutes.GET("/:id", func(c *gin.Context) { controllers.GetFormResponse(c, db) })
		formRoutes.PUT("/:id", func(c *gin.Context) { controllers.UpdateFormResponse(c, db) })
		formRoutes.DELETE("/:id", func(c *gin.Context) { controllers.DeleteFormResponse(c, db) })
	}
}
