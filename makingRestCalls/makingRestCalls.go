package main

import (
	"io/ioutil"
	"net/http"
)

func main() {

	http.HandleFunc("/employees", GetEmployeeData)
	http.ListenAndServe(":4000", nil)
}

func GetEmployeeData(w http.ResponseWriter, r *http.Request) {
	response, _ := http.Get("http://5c055de56b84ee00137d25a0.mockapi.io/api/v1/employees")
	data, _ := ioutil.ReadAll(response.Body)
	//jsonData := string(data)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(data)
}
