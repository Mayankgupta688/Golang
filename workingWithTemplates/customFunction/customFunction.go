package main

import (
	"html/template"
	"os"
)

type Employee struct {
	Name        string
	Designation string
}

func (emp Employee) ShowMessage() string {
	return "Ëmployee Name is: " + emp.Name
}

func main() {
	newEmployee := Employee{"Mayank", "Developer"}

	templateString := "{{ .ShowMessage }}"
	newTemplate, _ := template.New("customTemplate").Parse(templateString)
	newTemplate.Execute(os.Stdout, newEmployee)
}
