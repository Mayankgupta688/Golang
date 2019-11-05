package main

import (
	"fmt"
	"net/http"
	"path"
	"text/template"
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
	http.ListenAndServe(":8000", nil)
}

func ShowTemplateWithData(w http.ResponseWriter, r *http.Request) {

	fp := path.Join("template", "index.html")

	fmt.Println(fp)
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(w, nil)
}
