package routes

import (
    "github.com/Makeyabe/Home_Backend/controllers"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// SetupBookingRoutes - ตั้งค่าเส้นทางการจอง
func SetupBookingRoutes(router *gin.Engine, db *gorm.DB) {
    bookingController := &controllers.BookingController{DB: db}

    bookings := router.Group("/api/bookings")
    {
        bookings.POST("/", bookingController.CreateBooking) // สร้างการจอง
        bookings.GET("/", bookingController.GetBookings)    // ดึงข้อมูลการจองทั้งหมด
        // bookings.DELETE("/:id", bookingController.DeleteBooking) // ยกเลิกการจอง
    }
}
