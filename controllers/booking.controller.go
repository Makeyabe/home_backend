package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BookingController struct
type BookingController struct {
	DB *gorm.DB
}

func NewBookingController(db *gorm.DB) *BookingController {
	return &BookingController{DB: db}
}

// CreateBooking - สร้างการจองใหม่
func (bc *BookingController) CreateBooking(c *gin.Context) {
	var req model.BookingRequest

	// Bind and parse the JSON payload
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking data"})
		return
	}

	// Parse the booking date
	bookingDate, err := time.Parse(time.RFC3339, req.BookingDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format or date out of range"})
		return
	}
	bookingDate = bookingDate.Truncate(24 * time.Hour) // Keep only the date part

	// Start a transaction
	tx := bc.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if there are already 3 bookings on this date
	var bookingCount int64
	if err := tx.Model(&model.Booking{}).
		Where("DATE(booking_date) = ?", bookingDate).
		Count(&bookingCount).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check booking count"})
		return
	}

	if bookingCount >= 3 {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": "Booking limit reached for this date"})
		return
	}

	// Check if the student already has any booking record (ignoring the date)
	var existingBooking model.Booking
	// err = tx.Where("stu_id = ?", req.StuID).First(&existingBooking).Error

	stuID, err := strconv.Atoi(req.StuID) // แปลง req.StuID จาก string เป็น int
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID format"})
		return
	}

	err = tx.Where("stu_id = ?", stuID).First(&existingBooking).Error
	if err != nil {
		// handle error เช่น แจ้งข้อผิดพลาดกลับไปที่ client
	}

	if err == nil {
		// Existing booking found, update it with the new date
		existingBooking.BookingDate = bookingDate
		existingBooking.StuName = req.StuName
		if err := tx.Save(&existingBooking).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update booking"})
			return
		}
	} else if err == gorm.ErrRecordNotFound {
		// No existing booking found, create a new one
		newBooking := model.Booking{
			StuId:       req.StuID,
			StuName:     req.StuName,
			BookingDate: bookingDate,
		}
		if err := tx.Create(&newBooking).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"})
			return
		}
	} else {
		// Some other error occurred
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking existing booking"})
		return
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to finalize booking"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Booking processed successfully"})
}

// GetBookings - ดึงข้อมูลการจองทั้งหมด
func (bc *BookingController) GetBookings(c *gin.Context) {
	var bookings []model.Booking
	if err := bc.DB.Find(&bookings).Error; err != nil {
		log.Printf("Error retrieving bookings: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

// DeleteBooking - ยกเลิกการจอง
func (bc *BookingController) DeleteBooking(c *gin.Context) {
	id := c.Param("id")
	bookingID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking ID"})
		return
	}

	if err := bc.DB.Delete(&model.Booking{}, bookingID).Error; err != nil {
		log.Printf("Error deleting booking with ID %d: %v", bookingID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete booking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Booking deleted successfully"})
}
