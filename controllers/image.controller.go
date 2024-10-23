package controllers

import (
	"net/http"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB // Database instance

// CreateImage สร้างภาพใหม่
func CreateImage(c *gin.Context) {
	var image model.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&image)
	c.JSON(http.StatusCreated, image)
}

// GetImages รับข้อมูลภาพทั้งหมด
func GetImages(c *gin.Context) {
	var images []model.Image
	db.Find(&images)
	c.JSON(http.StatusOK, images)
}

// GetImage รับข้อมูลภาพตาม ID
func GetImage(c *gin.Context) {
	var image model.Image
	id := c.Param("id")
	if err := db.First(&image, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	c.JSON(http.StatusOK, image)
}

// UpdateImage อัปเดตข้อมูลภาพตาม ID
func UpdateImage(c *gin.Context) {
	var image model.Image
	id := c.Param("id")
	if err := db.First(&image, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&image)
	c.JSON(http.StatusOK, image)
}

// DeleteImage ลบภาพตาม ID
func DeleteImage(c *gin.Context) {
	id := c.Param("id")
	if err := db.Delete(&model.Image{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// SetDB ตั้งค่า db สำหรับคอนโทรลเลอร์
func SetDB(database *gorm.DB) {
	db = database
}
