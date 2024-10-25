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
	ID        uint               `gorm:"primaryKey"`
	TeacherID string             `gorm:"not null"` // Foreign key to Teacher
	StudentID string             `gorm:"not null"` // Foreign key to Student
	Term      string             `gorm:"not null"`
	Names     []NameEntry        `gorm:"type:jsonb;serializer:json"`      // Store names as JSONB for structured data
	Sections  []ResponseSection  `gorm:"foreignKey:ResponseFormID;references:ID"` // One-to-many relationship with ResponseSection
	CreatedAt time.Time          `json:"-"`
	UpdatedAt time.Time          `json:"-"`
	DeletedAt gorm.DeletedAt     `json:"-"`
}

// ResponseSection represents sections within the form response
type ResponseSection struct {
	ID             uint            `gorm:"primaryKey"`
	ResponseFormID uint            `gorm:"not null"`             // Foreign key to ResponseForm
	SectionID      uint            `gorm:"not null"`             // Section ID (unique within a form)
	Title          string          `gorm:"not null"`             // Title of the section (e.g., "Housing", "Environment")
	Fields         []ResponseField `gorm:"foreignKey:SectionID"` // One-to-many relationship with ResponseField
	CreatedAt      time.Time       `json:"-"`
	UpdatedAt      time.Time       `json:"-"`
	DeletedAt      gorm.DeletedAt  `json:"-"`
}

// ResponseField represents individual fields in a section
type ResponseField struct {
	ID        uint           `gorm:"primaryKey"`
	SectionID uint           `gorm:"not null"`  // Foreign key to ResponseSection
	FieldID   uint           `gorm:"not null"`  // Unique ID for the field within the section
	Value     string         `gorm:"type:text"` // Value that the user submitted for this field
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
