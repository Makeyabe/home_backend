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
	TeacherID uint                  `gorm:"not null"`                                 // Foreign key to Teacher
	Teacher   Teacher               `gorm:"foreignKey:TeacherID;references:Username"` // Foreign key to Teacher's Username
	StudentID uint                  `gorm:"not null"`                                 // Foreign key to Student
	Student   Student               `gorm:"foreignKey:StudentID;references:StuId"`    // Foreign key to Student's StuId
	Term      string                `gorm:"not null"`
	Names     []NameEntry           `gorm:"type:jsonb"`            // Store names as JSONB for structured data
	Sections  []FormResponseSection `gorm:"foreignKey:ResponseID"` // One-to-many relationship with FormResponseSection
	CreatedAt time.Time             `json:"-"`
	UpdatedAt time.Time             `json:"-"`
	DeletedAt gorm.DeletedAt        `json:"-"`
}

// FormResponseSection represents the sections submitted in the form
type FormResponseSection struct {
	ID         uint                `gorm:"primaryKey"`
	ResponseID uint                `gorm:"not null"`             // Foreign key to FormResponse (i.e., which form this section belongs to)
	SectionID  uint                `gorm:"not null"`             // Section ID (unique within a form)
	Title      string              `gorm:"not null"`             // Title of the section (e.g., "Housing", "Environment")
	Fields     []FormResponseField `gorm:"foreignKey:SectionID"` // One-to-many relationship with FormResponseField
	CreatedAt  time.Time           `json:"-"`
	UpdatedAt  time.Time           `json:"-"`
	DeletedAt  gorm.DeletedAt      `json:"-"`
}

type FormResponseField struct {
	ID        uint           `gorm:"primaryKey"`
	SectionID uint           `gorm:"not null"`  // Foreign key to FormResponseSection
	FieldID   uint           `gorm:"not null"`  // Unique ID for the field within the section
	Value     string         `gorm:"type:text"` // Value that the user submitted for this field (e.g., text, answer)
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}