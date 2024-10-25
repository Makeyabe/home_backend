package controllers

import (
	"log"
	"net/http"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentController struct {
	DB *gorm.DB
}

func NewStudentController(db *gorm.DB) *StudentController {
	return &StudentController{DB: db}
}

type StudentLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (sc *StudentController) StudentLogin(c *gin.Context) {
	var input StudentLoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var student model.Student // สมมุติว่ามี model ที่ชื่อว่า Student
	if err := sc.DB.Where("username = ?", input.Username).First(&student).Error; err != nil {
		c.JSON(400, gin.H{"message": "Student not found"})
		return
	}

	if student.Password != input.Password {
		c.JSON(400, gin.H{"message": "Incorrect password"})
		return
	}

	// หากต้องการให้ลบส่วนเกี่ยวกับ token ที่นี่ เช่น
	c.JSON(200, gin.H{"message": "Login successful"})
}

func (sc *StudentController) GetStudentData(c *gin.Context) {
	var students []*model.Student

	tx := sc.DB.Begin()
	if tx.Error != nil {
		c.JSON(500, gin.H{"error": "Failed to start transaction"})
		return
	}

	if err := tx.Find(&students).Error; err != nil {
		tx.Rollback() // Rollback on error
		c.JSON(500, gin.H{"error": "Failed to retrieve students"})
		return
	}

	studentUsernames := make([]string, len(students))
	for i, student := range students {
		studentUsernames[i] = student.Username
	}

	var responseForms []*model.ResponseForm
	if err := tx.Where("student_id IN ? AND term IN ?", studentUsernames, []string{"1", "2"}).Find(&responseForms).Error; err != nil {
		tx.Rollback() // Rollback on error
		c.JSON(500, gin.H{"error": "Failed to retrieve response forms"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to commit transaction"})
		return
	}

	responseMap := make(map[string]map[string]bool)
	for _, responseForm := range responseForms {
		if responseMap[responseForm.StudentID] == nil {
			responseMap[responseForm.StudentID] = make(map[string]bool)
		}
		responseMap[responseForm.StudentID][responseForm.Term] = true
	}

	for _, student := range students {
		student.FirstVisit = responseMap[student.Username]["1"]
		student.SecondVisit = responseMap[student.Username]["2"]
	}

	c.JSON(200, students)
}


func (sc *StudentController) GetStudentByID(c *gin.Context) {
	// ดึงค่า ID จากพารามิเตอร์ใน URL
	studentID := c.Param("id")

	var student model.Student

	// ดึงข้อมูลนักเรียนจากฐานข้อมูลโดยใช้ ID
	if err := sc.DB.Where("username = ?", studentID).First(&student).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	// ส่งข้อมูลนักเรียนกลับไปยัง client
	c.JSON(200, student)
}

func (sc *StudentController) UpdateStudent(ctx *gin.Context) {
	var student model.Student // Variable to hold the updated data
	id := ctx.Param("id")     // Get the student ID from the URL
	log.Println("Student", id)

	// Bind the JSON body to the student struct
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use GORM to update the student in the database
	if err := sc.DB.Model(&student).Where("username = ?", id).Updates(student).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update student"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Student updated successfully"})
}

func (sc *StudentController) SaveStudent(c *gin.Context) {
	var student model.Student

	// รับข้อมูล JSON ที่ส่งมาจาก client และ bind เข้ากับ student struct
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// ตรวจสอบว่ามี username หรือ stu_id ที่ซ้ำกันในฐานข้อมูลหรือไม่
	var existingStudent model.Student
	if err := sc.DB.Where("username = ? OR stu_id = ?", student.Username, student.StuId).First(&existingStudent).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username or Student ID already exists"})
		return
	}

	// บันทึกข้อมูลลงในฐานข้อมูล
	if err := sc.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Student saved successfully", "data": student})
}
