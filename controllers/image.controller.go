package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ImageController struct {
	DB *gorm.DB // Database instance
}

func NewImageController(db *gorm.DB) *ImageController {
	return &ImageController{DB: db}
}

const (
	ImageSaveDir  = "./uploads/images" // Directory where images are saved
	MaxFileSize   = 10 * 1024 * 1024   // 10 MB in bytes
	JPG           = "jpg"
	PNG           = "png"
)

// isValidImage checks if the file extension is either .jpg or .png
func isValidImage(filename string) bool {
	ext := filepath.Ext(filename)
	return ext == "."+JPG || ext == "."+PNG
}

// CreateImage handles saving an uploaded image and storing metadata in the database
func (IC *ImageController) CreateImage(c *gin.Context) {
	// Parse `stu_id` from form-data
	stuIDStr := c.PostForm("stu_id")
	stuID, err := strconv.Atoi(stuIDStr) // Convert stu_id from string to int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stu_id must be an integer"})
		return
	}

	// Retrieve file from form-data with the key "image"
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image is uploaded"})
		return
	}

	// Check if file size is within the 10 MB limit
	if file.Size > MaxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the 10MB limit"})
		return
	}

	// Check if file extension is .jpg or .png
	if !isValidImage(file.Filename) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Only .jpg and .png file types are allowed"})
		return
	}

	// Create a unique filename using stu_id and current timestamp
	ext := filepath.Ext(file.Filename) // Get file extension
	filename := fmt.Sprintf("stu_%d_%d%s", stuID, time.Now().Unix(), ext)
	filePath := filepath.Join(ImageSaveDir, filename)

	// Save the uploaded file to the specified path
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the image"})
		return
	}

	// Create a new Image record with the unique filename
	image := model.Image{
		StuId:     stuID,
		Imagepath: filename, // Store only the filename, not the full path
	}

	// Save to the database
	if err := IC.DB.Create(&image).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image data to database"})
		return
	}

	// Return response with saved image details
	c.JSON(http.StatusCreated, gin.H{
		"message": "Image uploaded successfully",
		"image":   image,
	})
}
