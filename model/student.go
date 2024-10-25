package model

type Student struct {
	ID            int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username      string `gorm:"type:varchar(255)" json:"username"`
	Password      string `gorm:"type:varchar(255)" json:"password"` // รหัสผ่าน
	Name          string `gorm:"type:varchar(255)" json:"name"`     // ชื่อ
	Nickname      string `gorm:"type:varchar(255)" json:"nickname"` // ชื่อเล่น
	Idcard        string `gorm:"type:varchar(255)" json:"idcard"`
	StuId         int    `gorm:"type:integer;unique" json:"stu_id"` // StuId is an integer
	StuPhone      string `gorm:"type:varchar(100)" json:"stu_phone"`
	StuClass      string `gorm:"type:varchar(10)" json:"stu_class"`
	StuBirthDate  string `gorm:"type:varchar(255)" json:"stu_birth_date"`
	Address       string `gorm:"type:varchar(255)" json:"address"`
	Distance      string `gorm:"type:varchar(100)" json:"distance"`
	Transport     string `gorm:"type:varchar(255)" json:"transport"`
	Skills        string `gorm:"type:varchar(255)" json:"skills"`
	FatherName    string `gorm:"type:varchar(255)" json:"father_name"`
	FatherJob     string `gorm:"type:varchar(255)" json:"father_job"`
	FatherPhone   string `gorm:"type:varchar(255)" json:"father_phone"`
	FatherSalary  int    `gorm:"type:integer" json:"father_salary"`
	FatherEdu     string `gorm:"type:varchar(255)" json:"father_edu"`
	MotherName    string `gorm:"type:varchar(255)" json:"mother_name"`
	MotherJob     string `gorm:"type:varchar(255)" json:"mother_job"`
	MotherPhone   string `gorm:"type:varchar(255)" json:"mother_phone"`
	MotherSalary  int    `gorm:"type:integer" json:"mother_salary"`
	MotherEdu     string `gorm:"type:varchar(255)" json:"mother_edu"`
	ParentName    string `gorm:"type:varchar(255)" json:"parent_name"`
	Relation      string `gorm:"type:varchar(255)" json:"relation"`
	ParentPhone   string `gorm:"type:varchar(255)" json:"parent_phone"`
	ParentAddress string `gorm:"type:varchar(255)" json:"parent_address"`
	PStatus       string `gorm:"type:varchar(255)" json:"p_status"`
	LivesWith     string `gorm:"type:varchar(255)" json:"lives_with"`
	FamCount      string `gorm:"type:varchar(255)" json:"fam_count"`
	SibStudy      string `gorm:"type:varchar(255)" json:"sib_study"`
	EmpCount      string `gorm:"type:varchar(255)" json:"emp_count"`
	UnempCount    string `gorm:"type:varchar(255)" json:"unemp_count"`
	MapUrl        string `gorm:"type:text" json:"map_url"`

	FirstVisit  bool `json:"first_visit"`
	SecondVisit bool `json:"second_visit"`
}

type SignInStudent struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GetStudentData struct {
	Name          string `json:"name"`
	Nickname      string `json:"nickname"`
	StuId         string `json:"stuid"`
	StuPhone      string `json:"stuphone"`
	StuClass      string `json:"stuclass"`
	StuBirthDate  string `json:"stubirthdate"`
	Address       string `json:"address"`
	Distance      string `json:"distance"`
	Transport     string `json:"transport"`
	Skills        string `json:"skills"`
	FatherName    string `json:"fathername"`
	FatherJob     string `json:"fatherjob"`
	FatherPhone   string `json:"fatherphone"`
	FatherSalary  int    `json:"fathersalary"`
	FatherEdu     string `json:"fatheredu"`
	MotherName    string `json:"mothername"`
	MotherJob     string `json:"motherjob"`
	MotherPhone   string `json:"motherphone"`
	MotherSalary  int    `json:"mothersalary"`
	MotherEdu     string `json:"motheredu"`
	ParentName    string `json:"parentname"`
	Relation      string `json:"relation"`
	ParentPhone   string `json:"parentphone"`
	ParentAddress string `json:"parentaddress"`
	PStatus       string `json:"pstatus"`
	LivesWith     string `json:"livewith"`
	FamCount      string `json:"famcount"`
	SibStudy      string `json:"sibstudy"`
	EmpCount      string `json:"empcount"`
	UnempCount    string `json:"unempcount"`
	MapUrl        string `json:"mapurl"`
}
