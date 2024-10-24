package model

import (
	"time"

	"gorm.io/gorm"
)

type NameEntry struct {
	Role string `json:"role"`
	Name string `json:"name"`
}

type FormResponse struct {
	ID        uint                  `gorm:"primaryKey"`
	TeacherID uint                  `gorm:"not null"`                                     // Foreign key to Teacher
	Teacher   Teacher               `gorm:"foreignKey:TeacherID;references:Username"`     // Foreign key to Teacher's Username
	StudentID uint                  `gorm:"not null"`                                     // Foreign key to Student
	Student   Student               `gorm:"foreignKey:StudentID;references:StuId"`        // Foreign key to Student's StuId
	Term      string                `gorm:"not null"`
	Sections  []FormResponseSection `gorm:"foreignKey:ResponseID"`                        // One-to-many relationship with FormResponseSection
	Names     []NameEntry           `gorm:"type:jsonb"`                                   // Store names as jsonb (for complex data)
	CreatedAt time.Time             `json:"-"`                                            // Exclude CreatedAt from JSON
	UpdatedAt time.Time             `json:"-"`                                            // Exclude UpdatedAt from JSON
	DeletedAt gorm.DeletedAt        `json:"-"`                                            // Exclude DeletedAt from JSON
}

// FormResponseSection represents the sections submitted in the form
type FormResponseSection struct {
	ID         uint                `gorm:"primaryKey"`
	ResponseID uint                `gorm:"not null"`             // Foreign key to FormResponse
	SectionID  uint                `gorm:"not null"`             // Section ID
	Title      string              `gorm:"not null"`             // Title of the section
	Fields     []FormResponseField `gorm:"foreignKey:SectionID"` // One-to-many relationship with FormResponseField
	CreatedAt  time.Time           `json:"-"`
	UpdatedAt  time.Time           `json:"-"`
	DeletedAt  gorm.DeletedAt      `json:"-"`
}

// FormResponseField represents the response to each field in a form
type FormResponseField struct {
	ID        uint           `gorm:"primaryKey"`
	SectionID uint           `gorm:"not null"`  // Foreign key to FormResponseSection
	FieldID   uint           `gorm:"not null"`  // Foreign key to FormField
	Value     string         `gorm:"type:text"` // Value that the user submitted for this field
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
