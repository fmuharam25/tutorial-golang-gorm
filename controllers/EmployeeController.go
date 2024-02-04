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

func UpdateEmployee(id uint, name string, dept *Department) (*Employee, error) {
	emp := &Employee{ID: id, Name: name, DepartmentID: dept.ID}
	err := db.Save(emp).Error
	return emp, err
}

func DeleteEmployee(id uint) error {
	emp := &Employee{ID: id}
	return db.Delete(emp).Error
}
