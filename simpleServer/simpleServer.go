package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandleIndex)
	http.ListenAndServe(":3000", nil)
}

type Employee struct {
	name string
	age  int
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	newEmp := Employee{"Mayank", 30}
	fmt.Fprintf(w, "Hello "+newEmp.name)

}
