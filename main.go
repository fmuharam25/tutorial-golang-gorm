package main

import (
	"fmt"

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
	employeMove, err := c.UpdateEmployee(employee1.ID, "John Doe", department2)

	fmt.Println("===Task 2===")
	fmt.Println("Department :", department2.Name)
	fmt.Println("Move Employee :", employeMove.Name)

	//Task 3

	// Create new department
	departments := make([]Department, 3)
	for i := 1; i <= 3; i++ {
		departments = append(departments, Department{Name: fmt.Sprintf("Department %d", i)})
	}

	employees := make([]Employee, 5)
	for i := 1; i <= 5; i++ {
		employees = append(employees, Employee{Name: fmt.Sprintf("Employee %d", i), Department: departments[i%3]})
	}

	db.CreateInBatches(departments, 3)
	db.CreateInBatches(employees, 3)

	fmt.Println("===Task 3===")
	fmt.Println("Department :", departments)
	fmt.Println("Move Employee :", employees)

}
