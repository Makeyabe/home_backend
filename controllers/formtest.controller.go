package controllers

import (
	"log"
	"net/http"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FormController struct {
	DB *gorm.DB
}

// NewFormController สร้างตัวควบคุมฟอร์มใหม่
func NewFormController(db *gorm.DB) *FormController {
	return &FormController{DB: db}
}

// GetForms ดึงข้อมูลฟอร์มทั้งหมด
func (fc *FormController) GetForms(ctx *gin.Context) {
	var forms []model.Form // Slice to hold multiple forms

	// Preload associated Sections and Fields together
	if err := fc.DB.Preload("Sections.Fields").Find(&forms).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Println(forms)

	// Send the data back as a JSON response
	ctx.JSON(http.StatusOK, forms)
}

// GetForm ดึงข้อมูลฟอร์มตาม ID
func (fc *FormController) GetForm(ctx *gin.Context) {
	var form model.Form
	formID := ctx.Param("id") // ดึง ID ของฟอร์มจาก URL

	// Preload เพื่อดึงข้อมูล Sections และ Fields ของฟอร์ม
	if err := fc.DB.Preload("Sections.Fields").First(&form, formID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		return
	}

	// ส่งข้อมูลฟอร์มกลับไป
	ctx.JSON(http.StatusOK, form)
}

// CreateForm สร้างฟอร์มใหม่
func (fc *FormController) CreateForm(ctx *gin.Context) {
	var form model.Form
	// ดึงข้อมูลฟอร์มจาก request body
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// สร้างฟอร์มในฐานข้อมูล
	if err := fc.DB.Create(&form).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// ส่งข้อมูลฟอร์มที่ถูกสร้างกลับไป
	ctx.JSON(http.StatusCreated, form)
}

// UpdateForm อัปเดตฟอร์มที่มีอยู่
func (fc *FormController) UpdateForm(ctx *gin.Context) {
	var form model.Form
	formID := ctx.Param("id") // ดึง ID ของฟอร์มจาก URL

	// ค้นหาฟอร์มในฐานข้อมูล
	if err := fc.DB.First(&form, formID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		return
	}

	// ดึงข้อมูลใหม่จาก request body
	if err := ctx.ShouldBindJSON(&form); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตฟอร์มในฐานข้อมูล
	if err := fc.DB.Save(&form).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งข้อมูลฟอร์มที่ถูกอัปเดตกลับไป
	ctx.JSON(http.StatusOK, form)
}

// DeleteForm ลบฟอร์ม
func (fc *FormController) DeleteForm(ctx *gin.Context) {
	var form model.Form
	formID := ctx.Param("id") // ดึง ID ของฟอร์มจาก URL

	// ค้นหาฟอร์มในฐานข้อมูล
	if err := fc.DB.First(&form, formID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Form not found"})
		return
	}

	// ลบฟอร์ม
	if err := fc.DB.Delete(&form).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ส่งการตอบกลับว่าได้ลบสำเร็จ
	ctx.JSON(http.StatusOK, gin.H{"message": "Form deleted successfully"})
}
