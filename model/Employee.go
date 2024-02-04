package model

type Employee struct {
	ID           uint   `gorm:"primarykey"`
	Name         string `gorm:"type:varchar(100)"`
	DepartmentID uint
	Department   Department `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
