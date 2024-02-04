package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	c "go_gorm/controllers"
	. "go_gorm/model"

	"go_gorm/database"
)

var db = database.Connect()

func main() {
	//Migrate db
	db.AutoMigrate(&Department{}, &Employee{})

	//Task 1
	// Create department
	department, err := c.CreateDepartment("IT")

	// Create employees
	employee1, err := c.CreateEmployee("John Doe", department)
	employee2, err := c.CreateEmployee("Doe Doe", department)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("=== Task 1===")
	fmt.Println("Department :", department.Name)
	fmt.Println("Employee1 :", employee1.Name)
	fmt.Println("Employee2 :", employee2.Name)

	//Task 2

	// Create new department
	department2, err := c.CreateDepartment("HR")

	// Move employees
	employeMove, err := c.UpdateEmployee(employee1.ID, "John Doe", department2.ID)

	fmt.Println("===Task 2===")
	fmt.Println("Department :", department2.Name)
	fmt.Println("Move Employee :", employeMove.Name)

	//Task 3

	// Append new department
	departments := []Department{}
	for i := 1; i <= 3; i++ {
		departments = append(departments, Department{Name: "New Batch Department" + strconv.Itoa(i)})
	}
	//Create batch
	db.CreateInBatches(departments, 3)

	// Append employe and add to random department
	employees := []Employee{}
	for i := 1; i <= 5; i++ {
		employees = append(employees, Employee{Name: "New Batch Employee" + strconv.Itoa(i), DepartmentID: departments[i%2].ID})
	}
	db.CreateInBatches(employees, 5)

	//Get random
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	idep := r.Intn(len(departments))

	idx := []uint{}
	for i := 1; i <= 3; i++ {
		idx = append(idx, employees[r.Intn(len(employees))].ID)
	}
	// Update employe to
	ids := c.UpdateBatchEmployee(idx, departments[idep].ID)

	fmt.Println("===Task 3===")
	fmt.Println("New Department :", departments)
	fmt.Println("Add Employee To Exists Department :", employees)
	fmt.Println("Update Employee Ids :", ids)

}
