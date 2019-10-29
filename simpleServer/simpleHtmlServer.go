package main

import (
	"net/http"
	"path"
	"text/template"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	http.HandleFunc("/", ShowTemplateWithData)
	http.ListenAndServe(":3000", nil)
}

func ShowTemplateWithData(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Mayank Gupta"}
	fp := path.Join("templates", "sample.html")
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(w, book)
}
