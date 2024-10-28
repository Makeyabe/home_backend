package controllers

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strings"

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

	// Fetch all students
	if err := sc.DB.Find(&students).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve students"})
		return
	}

	// Populate visit data for each student
	if err := sc.populateVisitData(students); err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve visit information"})
		return
	}

	// Send response with students data
	c.JSON(200, students)
}

func (sc *StudentController) GetStudentByID(c *gin.Context) {
	// Retrieve student ID from URL parameter
	studentID := c.Param("id")
	var student model.Student

	// Fetch student by ID
	if err := sc.DB.Where("username = ?", studentID).First(&student).Error; err != nil {
		c.JSON(404, gin.H{"error": "Student not found"})
		return
	}

	// Populate visit data for the single student
	if err := sc.populateVisitData([]*model.Student{&student}); err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve visit information"})
		return
	}

	// Send response with student data
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

func (sc *StudentController) populateVisitData(students []*model.Student) error {
	// Extract student usernames to fetch visit records
	studentUsernames := make([]string, len(students))
	for i, student := range students {
		studentUsernames[i] = student.Username
	}

	// Fetch ResponseForm data for terms 1 and 2 for the students
	var responseForms []*model.ResponseForm
	if err := sc.DB.Where("student_id IN ? AND term IN ?", studentUsernames, []string{"1", "2"}).Find(&responseForms).Error; err != nil {
		return err
	}

	// Map to store visit data by student and term
	responseMap := make(map[string]map[string]bool)
	for _, responseForm := range responseForms {
		if responseMap[responseForm.StudentID] == nil {
			responseMap[responseForm.StudentID] = make(map[string]bool)
		}
		responseMap[responseForm.StudentID][responseForm.Term] = true
	}

	// Update each student's FirstVisit and SecondVisit based on responseMap
	for _, student := range students {
		student.FirstVisit = responseMap[student.Username]["1"]
		student.SecondVisit = responseMap[student.Username]["2"]
	}

	return nil
}

type SectionReport struct {
	Level       string  `json:"level"`
	Count       int     `json:"count"`
	Percentage  float64 `json:"percentage"`
	Description string  `json:"description"`
}

type TermSummary struct {
	TotalStudents int                        `json:"total_students"`
	Sections      map[string][]SectionReport `json:"sections"`
}

type SummaryResponse struct {
	Terms    map[string]TermSummary `json:"terms"`
	Combined TermSummary            `json:"combined"`
}

func (sc *StudentController) SummaryReport(c *gin.Context) {
	// Fetch unique terms
	var terms []string
	if err := sc.DB.Model(&model.ResponseForm{}).
		Distinct("term").
		Pluck("term", &terms).Error; err != nil {
		log.Println("Error fetching terms:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูลเทอมได้"})
		return
	}

	// Descriptions for each level
	levelDescriptions := map[int]string{
		1: "ต้องการความช่วยเหลือเร่งด่วน",
		2: "ต้องการความช่วยเหลือ",
		3: "ต้องการความช่วยเหลือปานกลาง",
		4: "อยู่ในสภาพดี",
		5: "อยู่ในสภาพดีมาก",
	}

	// Initialize response structure
	response := SummaryResponse{
		Terms:    make(map[string]TermSummary),
		Combined: TermSummary{TotalStudents: 0, Sections: make(map[string][]SectionReport)},
	}

	combinedScoreCount := map[string]map[int]int{
		"บริบท ( Context – C )":                        {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
		"ครอบครัว ( Family – F )":                      {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
		"นักเรียน (Student – S )":                      {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
		"ความต้องการต่อโรงเรียน / อบจ. ( School – S )": {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
	}
	combinedTotalStudents := 0

	for _, term := range terms {
		// Count unique students for this term
		var countStudents int64
		if err := sc.DB.Model(&model.ResponseForm{}).
			Where("term = ?", term).
			Distinct("student_id").
			Count(&countStudents).Error; err != nil {
			log.Println("Error counting students for term", term, ":", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถนับจำนวนนักเรียนได้"})
			return
		}
		combinedTotalStudents += int(countStudents)

		// Fetch sections and fields for this term
		var sections []model.ResponseSection
		if err := sc.DB.Preload("Fields").
			Joins("JOIN response_forms ON response_forms.id = response_sections.response_form_id").
			Where("response_forms.term = ?", term).
			Find(&sections).Error; err != nil {
			log.Println("Error fetching sections for term", term, ":", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ไม่สามารถดึงข้อมูล section ได้"})
			return
		}

		// Initialize term-specific score count
		termScoreCount := map[string]map[int]int{
			"บริบท ( Context – C )":                        {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
			"ครอบครัว ( Family – F )":                      {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
			"นักเรียน (Student – S )":                      {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
			"ความต้องการต่อโรงเรียน / อบจ. ( School – S )": {1: 0, 2: 0, 3: 0, 4: 0, 5: 0},
		}

		// Track student average scores per section
		studentScores := make(map[string][]float64)
		for _, section := range sections {
			for _, field := range section.Fields {
				studentKey := fmt.Sprintf("%d-%s", section.ResponseFormID, section.Title)
				studentScores[studentKey] = append(studentScores[studentKey], float64(field.Score))
			}
		}

		// Calculate average and assign level
		for studentKey, scores := range studentScores {
			sectionTitle := strings.Split(studentKey, "-")[1]
			avgScore := calculateAverage(scores)
			roundedLevel := int(math.Round(avgScore))
			termScoreCount[sectionTitle][roundedLevel]++
			combinedScoreCount[sectionTitle][roundedLevel]++
		}

		// Build section reports for this term, sorted by level
		termSummary := TermSummary{
			TotalStudents: int(countStudents),
			Sections:      make(map[string][]SectionReport),
		}
		for sectionTitle, counts := range termScoreCount {
			var sectionReports []SectionReport
			for score := 1; score <= 5; score++ { // Loop from 1 to 5 to ensure order
				count := counts[score]
				percentage := (float64(count) / float64(countStudents)) * 100
				sectionReports = append(sectionReports, SectionReport{
					Level:       fmt.Sprintf("ระดับ %d", score),
					Count:       count,
					Percentage:  percentage,
					Description: levelDescriptions[score],
				})
			}
			termSummary.Sections[sectionTitle] = sectionReports
		}
		response.Terms[term] = termSummary
	}

	// Calculate combined summary across all terms, sorted by level
	response.Combined.TotalStudents = combinedTotalStudents
	for sectionTitle, counts := range combinedScoreCount {
		var sectionReports []SectionReport
		for score := 1; score <= 5; score++ { // Loop from 1 to 5 to ensure order
			count := counts[score]
			percentage := (float64(count) / float64(combinedTotalStudents)) * 100
			sectionReports = append(sectionReports, SectionReport{
				Level:       fmt.Sprintf("ระดับ %d", score),
				Count:       count,
				Percentage:  percentage,
				Description: levelDescriptions[score],
			})
		}
		response.Combined.Sections[sectionTitle] = sectionReports
	}

	c.JSON(http.StatusOK, response)
}

// Helper function to calculate the average of a float64 slice
func calculateAverage(scores []float64) float64 {
	var sum float64
	for _, score := range scores {
		sum += score
	}
	return sum / float64(len(scores))
}