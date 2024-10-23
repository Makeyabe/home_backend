package routes

import (
	"github.com/Makeyabe/Home_Backend/controllers"
	"github.com/gin-gonic/gin"
)

func FormResponseRoutes(router *gin.RouterGroup, FormResponseController *controllers.FormResponseController) {
	formRoutes := router.Group("/form-responses")
	{
		formRoutes.POST("/", FormResponseController.CreateFormResponse)
		formRoutes.GET("/:id", FormResponseController.GetFormResponse)   // ดึงฟอร์มตอบสนองตาม ID
		formRoutes.PUT("/:id", FormResponseController.UpdateFormResponse) // แก้ไขฟอร์มตอบสนองตาม ID
		formRoutes.DELETE("/:id", FormResponseController.DeleteFormResponse) // ลบฟอร์มตอบสนองตาม ID
	}
}
