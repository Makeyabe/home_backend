package routes

import (
    "github.com/Makeyabe/Home_Backend/controllers"
    "github.com/gin-gonic/gin"
)

func StudentRoutes(router *gin.RouterGroup, studentController *controllers.StudentController) {
    
    router.POST("/student/login", studentController.StudentLogin)
    router.GET("/students", studentController.GetStudentData) // เส้นทางใหม่สำหรับดึงข้อมูลนักเรียน
    router.GET("/student/:id", studentController.GetStudentByID) // เส้นทางใหม่สำหรับดึงข้อมูลนักเรียน
    router.GET("/student/summary", studentController.SummaryReport) // เส้นทางใหม่สำหรับดึงข้อมูลนักเรียน
    router.PUT("/student/:id", studentController.UpdateStudent)
    router.POST("/student/save", studentController.SaveStudent)
}
