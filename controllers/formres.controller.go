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
    var formSections model.FormResponseSection

	// // Bind JSON payload with formSections (array of sections)
	// if err := ctx.ShouldBindJSON(&formSections); err != nil {
	// 	log.Println("Failed to get body:", err)
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	// 	return
	// }

	// // Create a new FormResponse (assuming you have FormID, UserID, and Term from context or payload)
	// formResponse := model.FormResponse{
	// 	FormID:   1,  // Replace with actual value
	// 	UserID:   1,  // Replace with actual value
	// 	Term:     "2024 Term", // Replace with actual value from request if needed
	// 	Sections: formSections,
	// }

	// // Save the form response along with sections and fields
	// if err := db.Create(&formResponse).Error; err != nil {
	// 	log.Println("Failed to create form response:", err)
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create form response"})
	// 	return
	// } 

	ctx.JSON(http.StatusCreated, formSections)
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

func (frc *FormResponseController) CheckFormResponseExists(ctx *gin.Context) {
	var formResponse model.FormResponse

	// รับข้อมูลจาก JSON payload ที่ส่งมาจากหน้าบ้าน
	var input struct {
		FormID uint   `json:"form_id"`
		UserID uint   `json:"user_id"`
		Term   string `json:"term"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	// ค้นหาข้อมูลที่มี FormID, UserID และ Term ตรงกัน
	err := frc.DB.Where("form_id = ? AND user_id = ? AND term = ?", input.FormID, input.UserID, input.Term).First(&formResponse).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// ถ้าไม่พบข้อมูล
			ctx.JSON(http.StatusOK, gin.H{"exists": false, "message": "No form response found for this term"})
		} else {
			// ถ้าเกิดข้อผิดพลาดอื่นๆ
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check form response"})
		}
		return
	}

	// ถ้าพบข้อมูล
	ctx.JSON(http.StatusOK, gin.H{"exists": true, "message": "Form response found for this term", "data": formResponse})
}
