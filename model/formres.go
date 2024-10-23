package model

import (
	"time"

	"gorm.io/gorm"
)

// FormResponse represents the data submitted from a completed form
type FormResponse struct {
	ID        uint                   `gorm:"primaryKey"`
	FormID    uint                   `gorm:"not null"`                   // Foreign key to Form
	UserID    uint                   `gorm:"not null"`                   // User who submitted the form
	Fields    []FormResponseField    `gorm:"foreignKey:ResponseID"`      // One-to-many relationship with FormResponseField
	CreatedAt time.Time              `json:"-"`                          // Exclude CreatedAt from JSON
	UpdatedAt time.Time              `json:"-"`                          // Exclude UpdatedAt from JSON
	DeletedAt gorm.DeletedAt         `json:"-"`                          // Exclude DeletedAt from JSON
}

// FormResponseField represents the response to each field in a form
type FormResponseField struct {
	ID         uint           `gorm:"primaryKey"`
	ResponseID uint           `gorm:"not null"`                        // Foreign key to FormResponse
	FieldID    uint           `gorm:"not null"`                        // Foreign key to FormField
	Value      string         `gorm:"type:text"`                       // Value that the user submitted for this field
	CreatedAt  time.Time      `json:"-"`                               // Exclude CreatedAt from JSON
	UpdatedAt  time.Time      `json:"-"`                               // Exclude UpdatedAt from JSON
	DeletedAt  gorm.DeletedAt `json:"-"`                               // Exclude DeletedAt from JSON
}
