package controllers

import (
	. "go_gorm/model"
)

func CreateEmployee(name string, dept *Department) (*Employee, error) {
	emp := &Employee{Name: name, DepartmentID: dept.ID}
	err := db.Create(emp).Error
	return emp, err
}

func GetEmployee(id uint) (*Employee, error) {
	emp := &Employee{}
	err := db.First(emp, id).Error
	return emp, err
}

func UpdateEmployee(id uint, name string, deptId uint) (*Employee, error) {
	emp := &Employee{ID: id, Name: name, DepartmentID: deptId}
	err := db.Save(emp).Error
	return emp, err
}

func UpdateBatchEmployee(id []uint, deptId uint) []uint {

	db.Model(Employee{}).Where("id IN ?", id).Updates(Employee{DepartmentID: deptId})
	return id
}

func DeleteEmployee(id uint) error {
	emp := &Employee{ID: id}
	return db.Delete(emp).Error
}
