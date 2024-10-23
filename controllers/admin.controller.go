package controllers

// import (
// 	// "github.com/Makeyabe/Home_Backend/model"
// 	"github.com/gin-gonic/gin"
// 	"gorm.io/gorm"
// )

// type AdminController struct {
// 	DB *gorm.DB
// }

// func NewAdminController(DB *gorm.DB) AdminController {
// 	return AdminController{DB}
// }

// func (uc *AdminController) GetMe(ctx *gin.Context) {
// 	// ใช้ ctx.Get() เพื่อดึงข้อมูล currentUser และตรวจสอบว่ามีหรือไม่
// 	user, exists := ctx.Get("currentUser")
// 	if !exists {
// 		ctx.JSON(401, gin.H{"status": "fail", "message": "User not logged in"})
// 		return
// 	}

// 	// แปลงข้อมูลผู้ใช้ที่ดึงมาเป็น model.Admin
// 	currentUser, ok := user.(model.Admin)
// 	if !ok {
// 		ctx.JSON(500, gin.H{"status": "error", "message": "Failed to retrieve user data"})
// 		return
// 	}

// 	// สร้าง response สำหรับการส่งกลับ
// 	userResponse := &model.AdminResponse{
// 		ID:       currentUser.ID,
// 		Username: currentUser.Username,
// 		Role:     currentUser.Role,
// 	}

// 	// ส่ง response กลับไปยัง client
// 	ctx.JSON(200, gin.H{
// 		"status": "success",
// 		"data":   userResponse,
// 	})
// }
