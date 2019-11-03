package main

import (
	"Golang/privateStructPackage/privatepackage"
	"fmt"
)

func main() {
	newEmployee := privatepackage.OtherEmployeeDetails{"M", "A", "N"}

	fmt.Println(newEmployee.name)
}
