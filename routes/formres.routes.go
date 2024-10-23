package routes

import (
	"github.com/Makeyabe/Home_Backend/controllers"
	"github.com/gin-gonic/gin"
)

func FormResponseRoutes(router *gin.RouterGroup, FormResponseController *controllers.FormResponseController) {
	formRoutes := router.Group("/form-responses")
	{
		formRoutes.POST("/", FormResponseController.CreateFormResponse)
		// formRoutes.GET("/:id", func(c *gin.Context) { controllers.GetFormResponse(c, db) })
		// formRoutes.PUT("/:id", func(c *gin.Context) { controllers.UpdateFormResponse(c, db) })
		// formRoutes.DELETE("/:id", func(c *gin.Context) { controllers.DeleteFormResponse(c, db) })
	}
}
