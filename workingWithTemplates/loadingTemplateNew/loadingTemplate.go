package main

import (
	"html/template"
	"net/http"
)

func main() {
	customTemplates := populateTemplates()

	http.Handle("/scripts/", http.FileServer(http.Dir("./public")))
	http.Handle("/css/", http.FileServer(http.Dir("./public")))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := customTemplates.Lookup("index.html")
		t.Execute(w, nil)
	})

	http.ListenAndServe(":4000", nil)
}

func populateTemplates() *template.Template {
	customTemplates := template.New("customTemplates")
	template.Must(customTemplates.ParseGlob("templates/*.html"))
	return customTemplates
}
