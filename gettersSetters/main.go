package main

import "fmt"

type Person struct {
	name string
}

func (p Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func main() {
	

	user := new(Person)

	user.SetName("Anshul")

	fmt.Println(user.Name())
}