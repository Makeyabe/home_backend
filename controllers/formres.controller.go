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

	// Bind JSON payload กับตัวแปร formResponse
	if err := ctx.ShouldBindJSON(&formResponse); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึกลงฐานข้อมูล
	if err := db.Create(&formResponse).Error; err != nil {
		log.Println("Failed to create form response:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form response"})
		return
	}

	ctx.JSON(http.StatusCreated, formResponse)
}

// GetFormResponse รับข้อมูล FormResponse จากฐานข้อมูลตาม ID
func GetFormResponse(c *gin.Context, db *gorm.DB) {
	var formResponse model.FormResponse
	id := c.Param("id")

	// หา FormResponse ตาม ID
	if err := db.Preload("Fields").First(&formResponse, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Form response not found"})
		} else {
			log.Println("Failed to get form response:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get form response"})
		}
		return
	}

	c.JSON(http.StatusOK, formResponse)
}

// UpdateFormResponse แก้ไขข้อมูล FormResponse ตาม ID
func UpdateFormResponse(c *gin.Context, db *gorm.DB) {
	var formResponse model.FormResponse
	id := c.Param("id")

	// หา FormResponse ตาม ID
	if err := db.First(&formResponse, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Form response not found"})
		} else {
			log.Println("Failed to find form response:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find form response"})
		}
		return
	}

	// Bind JSON payload กับ formResponse
	if err := c.ShouldBindJSON(&formResponse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัพเดทลงฐานข้อมูล
	if err := db.Save(&formResponse).Error; err != nil {
		log.Println("Failed to update form response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update form response"})
		return
	}

	c.JSON(http.StatusOK, formResponse)
}

// DeleteFormResponse ลบ FormResponse ตาม ID
func DeleteFormResponse(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	// ลบข้อมูล
	if err := db.Delete(&model.FormResponse{}, id).Error; err != nil {
		log.Println("Failed to delete form response:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete form response"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Form response deleted"})
}
