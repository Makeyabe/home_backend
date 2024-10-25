package model

import (
	"time"

	"gorm.io/gorm"
)

type Mainsec struct {
	ID          uint           `gorm:"primaryKey"`
	Context     int            `gorm:"embedded"` // บริบท (Context)
	Family      int            `gorm:"embedded"` // ครอบครัว (Family)
	Student     int            `gorm:"embedded"` // นักเรียน (Student)
	SchoolNeeds int            `gorm:"embedded"` // ความต้องการต่อโรงเรียน/อบจ. (School Needs)
	CreatedAt   time.Time      `json:"-"`        // Exclude from JSON
	UpdatedAt   time.Time      `json:"-"`        // Exclude from JSON
	DeletedAt   gorm.DeletedAt `json:"-"`        // Exclude from JSON
}
