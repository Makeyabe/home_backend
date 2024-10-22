package model

import (
	"time"

	"gorm.io/gorm"
)

// Form represents the overall form structure (e.g., Home Visit Form)
// Form represents the overall form structure (e.g., Home Visit Form)
type Form struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(100);not null"` // Form name
	Sections  []FormSection  `gorm:"foreignKey:FormID"`          // One-to-many relationship with FormSection
	CreatedAt time.Time      `json:"-"`                          // Exclude CreatedAt from JSON
	UpdatedAt time.Time      `json:"-"`                          // Exclude UpdatedAt from JSON
	DeletedAt gorm.DeletedAt `json:"-"`                          // Exclude DeletedAt from JSON
}

// FormSection represents each section within the form (e.g., Housing, Environment)
// FormSection represents each section within the form (e.g., Housing, Environment)
type FormSection struct {
	ID             uint           `gorm:"primaryKey"`
	Title          string         `gorm:"type:varchar(100);not null"` // Title of the section (e.g., "Housing")
	FormID         uint           `gorm:"not null"`                   // Foreign key to Form
	SectionOrder   int            `gorm:"not null"`                   // Order of the section in the form
	Fields         []FormField    `gorm:"foreignKey:SectionID"`       // One-to-many relationship with FormField
	HasOptionsOnly bool           `gorm:"default:false"`              // True if section only has options without a label
	CreatedAt      time.Time      `json:"-"`                          // Exclude CreatedAt from JSON
	UpdatedAt      time.Time      `json:"-"`                          // Exclude UpdatedAt from JSON
	DeletedAt      gorm.DeletedAt `json:"-"`                          // Exclude DeletedAt from JSON
}

// FormField represents each question/field in a section
// FormField represents each question/field in a section
type FormField struct {
    ID         uint           `gorm:"primaryKey"`
    SectionID  uint           `gorm:"not null"`                                             // Foreign key to FormSection
    Label      string         `gorm:"type:varchar(255)"`                                    // Question/label for the field (e.g., "House Type")
    FieldType  string         `gorm:"type:varchar(50);not null"`                            // Field type (e.g., "text", "radio", "checkbox")
    Options    string         `gorm:"type:text"`                                            // JSON string of options for radio/checkbox fields (optional)
    FieldOrder int            `gorm:"not null"`                                             // Order of the field in the section
    FieldScore int            `gorm:"default:null"` // Score (must be between 1 and 5 and set default point as 1)
    CreatedAt  time.Time      `json:"-"`                                                    // Exclude CreatedAt from JSON
    UpdatedAt  time.Time      `json:"-"`                                                    // Exclude UpdatedAt from JSON
    DeletedAt  gorm.DeletedAt `json:"-"`                                                    // Exclude DeletedAt from JSON
}
