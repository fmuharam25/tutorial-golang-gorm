package controllers

import (
	. "go_gorm/model"
)

func CreateDepartment(name string) (*Department, error) {
	dept := &Department{Name: name}
	err := db.Create(dept).Error
	return dept, err
}

func GetDepartment(id uint) (*Department, error) {
	dept := &Department{}
	err := db.First(dept, id).Error
	return dept, err
}

func UpdateDepartment(id uint, name string) error {
	dept := &Department{ID: id, Name: name}
	return db.Save(dept).Error
}

func DeleteDepartment(id uint) error {
	dept := &Department{ID: id}
	return db.Delete(dept).Error
}
