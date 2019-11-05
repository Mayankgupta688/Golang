package main 

import "html/template"
import "net/http"

func main() {
	customTemplates := populateTemplates()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		templateFile := customTemplates.Lookup(requestedFile + ".html")
		templateFile.Execute(w, nil)
	})

	http.ListenAndServe(":4000", nil)
}


func populateTemplates() *template.Template {
	customTemplates := template.New("customTemplates")
	template.Must(customTemplates.ParseGlob("template" + "/*.html"))
	return customTemplates
}