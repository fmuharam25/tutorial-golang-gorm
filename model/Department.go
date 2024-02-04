package model

type Department struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"type:varchar(100);"` //add unique if u want
	Employees []Employee
}
