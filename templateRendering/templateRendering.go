package main

import (
	"net/http"
	"path"
	"text/template"
	"fmt"
)

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

type Todo struct {
	Title string
	Done  bool
}

func main() {
	http.HandleFunc("/", ShowTemplateWithData)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	http.ListenAndServe(":4000", nil)
}

func ShowTemplateWithData(w http.ResponseWriter, r *http.Request) {

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}
	fp := path.Join("template", "sample.html")

	fmt.Println(fp)
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(w, data)
}
