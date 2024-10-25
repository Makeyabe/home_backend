package model

type Term struct {
	TermID      string `json:"term_id"` // ระบุ ID ของเทอม เช่น "2024 Term"
	StuId       int    `gorm:"type:integer;unique" json:"stu_id"` //
	FirstVisit  bool   `json:"first_visit"`  // บันทึกสถานะการเยี่ยมครั้งแรก
	SecondVisit bool   `json:"second_visit"` // บันทึกสถานะการเยี่ยมครั้งที่สอง
}
