package model

import "time"

type Booking struct {
	ID          int       `gorm:"type:autoIncrement;primaryKey"`
	StuName     string    `gorm:"type:varchar(30)"`
	StuId       string    `gorm:"type:varchar(30)"`
	BookingDate time.Time `gorm:"type:timestamp"`
}

type BookingRequest struct {
	StuName     string 		`json:"name"`
	StuID       string    	`json:"studentId"`
	BookingDate string 		`json:"date"` // Use string temporarily
}
