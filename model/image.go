package model

type Image struct {
	ID uint  `gorm:"primaryKey;autoIncrement" json:"id"`
	StuId int `gorm:"type:integer"`
	Imagepath string `gorm:"type:varchar(255)"`
}