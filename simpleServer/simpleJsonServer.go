package main

import (
	"encoding/json"
	"net/http"
)

type Employee struct {
	Name string
	Age  int
}

func main() {
	http.HandleFunc("/", ShowBooks)
	http.ListenAndServe(":3000", nil)
}

func ShowBooks(w http.ResponseWriter, r *http.Request) {

	employeeList := []Employee{
		{"A", 30},
		{"B", 30},
		{"C", 30},
	}

	employeeList = append(employeeList, Employee{"D", 40})

	js, err := json.Marshal(employeeList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
