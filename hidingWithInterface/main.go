package main

import (
	"Golang/hidingWithInterface/emptype"
)

func main() {
	var interfacedOutput emptype.GetTeamSizeInterface

	interfacedOutput = emptype.Manager{"Mayank", "30", "50"}

	GetOutput(interfacedOutput)

	interfacedOutput = emptype.Lead{"Meha", "30", "100"}

	GetOutput(interfacedOutput)
}

func GetOutput(obj emptype.GetTeamSizeInterface) {
	obj.GetTeamSize()
}
