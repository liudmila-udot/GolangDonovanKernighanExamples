package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

func EmployeeById(id int) Employee {
	return Employee{
		ID:   1,
		Name: "Mila",
	}
}

func EmployeeByIdPointer(id int) *Employee {
	return &Employee{
		ID:   1,
		Name: "Mila",
	}
}

func main() {
	employee := EmployeeById(2)
	employeePointer := EmployeeByIdPointer(2)
	fmt.Println(employee.ID)
	fmt.Println(employeePointer.ID)

	id := employee.ID
	fmt.Println(id)
	id = 3
	fmt.Printf("After change in employee %d\n", employee.ID)

	idP := employeePointer.ID
	fmt.Println(idP)
	*&idP = 5
	fmt.Printf("After change in employeePointer %d\n", employeePointer.ID)

	// EmployeeById(2).ID = 3 Doesn't compile
	EmployeeByIdPointer(2).ID = 4
}
