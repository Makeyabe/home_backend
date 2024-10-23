package controllers

import (
	"log"
	"net/http"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FormResponseController struct {
	DB *gorm.DB
}

func NewFormResponseController(db *gorm.DB) *FormResponseController {
	return &FormResponseController{DB: db}
}

// CreateFormResponse รับข้อมูลการส่งฟอร์มจากผู้ใช้และบันทึกลงฐานข้อมูล
func (frc *FormResponseController) CreateFormResponse(ctx *gin.Context) {
	var formResponse model.FormResponse

	// // Bind JSON payload กับตัวแปร formResponse
	if err := ctx.ShouldBindJSON(&formResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": formResponse})
		return
	}

	// // บันทึกลงฐานข้อมูล
	if err := db.Create(&formResponse).Error; err != nil {
		log.Println("Failed to create form response:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form response"})
		return
	}

	ctx.JSON(http.StatusCreated, formResponse)
}

// GetFormResponse รับข้อมูล FormResponse จากฐานข้อมูลตาม ID
func (frc *FormResponseController) GetFormResponse(ctx *gin.Context) {
	var formResponse model.FormResponse
	id := ctx.Param("id")

	// หา FormResponse ตาม ID
	if err := frc.DB.Preload("Fields").First(&formResponse, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Form response not found"})
		} else {
			log.Println("Failed to get form response:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get form response"})
		}
		return
	}

	ctx.JSON(http.StatusOK, formResponse)
}

// UpdateFormResponse แก้ไขข้อมูล FormResponse ตาม ID
func (frc *FormResponseController) UpdateFormResponse(ctx *gin.Context) {
	var formResponse model.FormResponse
	id := ctx.Param("id")

	// หา FormResponse ตาม ID
	if err := frc.DB.First(&formResponse, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Form response not found"})
		} else {
			log.Println("Failed to find form response:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find form response"})
		}
		return
	}

	// Bind JSON payload กับ formResponse
	if err := ctx.ShouldBindJSON(&formResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัพเดทลงฐานข้อมูล
	if err := frc.DB.Save(&formResponse).Error; err != nil {
		log.Println("Failed to update form response:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update form response"})
		return
	}

	ctx.JSON(http.StatusOK, formResponse)
}

// DeleteFormResponse ลบ FormResponse ตาม ID
func (frc *FormResponseController) DeleteFormResponse(ctx *gin.Context) {
	id := ctx.Param("id")

	// ลบข้อมูล
	if err := frc.DB.Delete(&model.FormResponse{}, id).Error; err != nil {
		log.Println("Failed to delete form response:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete form response"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Form response deleted"})
}
