package main

import (
	"Golang/otherPrivateStructPackage/privatepackage"
	"fmt"
)

func main() {
	newEmployee := privatepackage.OtherEmployeeDetails{"M", "A", "N"}

	fmt.Println(newEmployee.name)
}
