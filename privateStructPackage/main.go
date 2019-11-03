package main

import (
	"fmt"
	"privatepackage"
)

func main() {
	newEmployee := privatepackage.EmployeeDetails{"M", "A", "N"}

	fmt.Println(newEmployee.Name)
}
