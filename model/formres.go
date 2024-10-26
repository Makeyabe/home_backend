package model

import (
	"time"

	"gorm.io/gorm"
)

// NameEntry represents the JSONB structured data for names
type NameEntry struct {
	Role string `json:"role"`
	Name string `json:"name"`
}

// ResponseForm represents the main form response
type ResponseForm struct {
	ID        uint              `gorm:"primaryKey" json:"id"`       // Primary key
	TeacherID string            `gorm:"not null" json:"teacher_id"` // Foreign key to Teacher
	StudentID string            `gorm:"not null" json:"student_id"` // Foreign key to Student
	Term      string            `gorm:"not null" json:"term"`
	Names     []NameEntry       `gorm:"type:jsonb;serializer:json" json:"names"`                 // Store names as JSONB for structured data
	Sections  []ResponseSection `gorm:"foreignKey:ResponseFormID;references:ID" json:"sections"` // One-to-many relationship with ResponseSection
	CreatedAt time.Time         `json:"created_at"`                                              // Send creation time
	UpdatedAt time.Time         `json:"updated_at"`                                              // Send update time
	DeletedAt gorm.DeletedAt    `json:"deleted_at,omitempty"`                                    // Send delete time if it exists
}

// ResponseSection represents sections within the form response
type ResponseSection struct {
	ID             uint            `gorm:"primaryKey" json:"id"`               // Primary key
	ResponseFormID uint            `gorm:"not null" json:"response_form_id"`   // Foreign key to ResponseForm
	SectionID      uint            `gorm:"not null" json:"section_id"`         // Section ID (unique within a form)
	Title          string          `gorm:"not null" json:"title"`              // Title of the section (e.g., "Housing", "Environment")
	Fields         []ResponseField `gorm:"foreignKey:SectionID" json:"fields"` // One-to-many relationship with ResponseField
	CreatedAt      time.Time       `json:"created_at"`                         // Send creation time
	UpdatedAt      time.Time       `json:"updated_at"`                         // Send update time
	DeletedAt      gorm.DeletedAt  `json:"deleted_at,omitempty"`               // Send delete time if it exists
}

// ResponseField represents individual fields in a section
type ResponseField struct {
	ID        uint           `gorm:"primaryKey" json:"id"`       // Primary key
	SectionID uint           `gorm:"not null" json:"section_id"` // Foreign key to ResponseSection
	FieldID   uint           `gorm:"not null" json:"field_id"`   // Unique ID for the field within the section
	Value     string         `gorm:"type:text" json:"value"`     // Value that the user submitted for this field
	Score     int            `gorm:"default:1" json:"score"`     // Score for this field (1-5)
	CreatedAt time.Time      `json:"created_at"`                 // Send creation time
	UpdatedAt time.Time      `json:"updated_at"`                 // Send update time
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`       // Send delete time if it exists
}
