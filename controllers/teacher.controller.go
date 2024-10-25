package controllers

import (
	"errors"
	"net/http"

	"github.com/Makeyabe/Home_Backend/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherController struct {
	DB *gorm.DB
}

func NewTeacherController(db *gorm.DB) *TeacherController {
	return &TeacherController{DB: db} // คืนค่า pointer
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (tc *TeacherController) Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	var teacher model.Teacher
	if err := tc.DB.Where("username = ?", input.Username).First(&teacher).Error; err != nil {
		c.JSON(400, gin.H{"error": "Teacher not found"})
		return
	}

	if teacher.Password != input.Password {
		c.JSON(400, gin.H{"error": "Incorrect password"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful"})
}

func (tc *TeacherController) GetStudentsByClass(c *gin.Context) {
    teacherID := c.Param("id")

    var teacher model.Teacher
	if err := tc.DB.Select("stu_class").Where("username = ?", teacherID).First(&teacher).Error; err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            c.JSON(http.StatusNotFound, gin.H{"error": "Teacher not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve teacher"})
        }
        return
    }
	
    var students []*model.Student // ใช้ pointer เพื่อประหยัด memory
    if err := tc.DB.Where("stu_class = ?", teacher.StuClass).Find(&students).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve students"})
        return
    }


	studentUsernames := make([]string, len(students))
    for i, student := range students {
        studentUsernames[i] = student.Username
    }

	var responseForms []*model.ResponseForm
	if err := tc.DB.Where("student_id IN ? AND term IN ?", studentUsernames, []string{"1", "2"}).Find(&responseForms).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve visit information"})
        return
    }

	responseMap := make(map[string]map[string]bool)
    for _, responseForm := range responseForms {
        if responseMap[responseForm.StudentID] == nil {
            responseMap[responseForm.StudentID] = make(map[string]bool)
        }
        responseMap[responseForm.StudentID][responseForm.Term] = true
    }

	 // Update students with visit information
	 for _, student := range students {
        student.FirstVisit = responseMap[student.Username]["1"]
        student.SecondVisit = responseMap[student.Username]["2"]
    } 

    // ส่งข้อมูลนักเรียนที่ตรงกับ stuClass กลับไปยัง client
    c.JSON(http.StatusOK, students)
}
