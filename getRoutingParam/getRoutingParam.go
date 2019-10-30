package main

import (
	"html/template"
	"net/http"
	"path"

	"github.com/julienschmidt/httprouter"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	r := httprouter.New()
	r.GET("/", GetIndexPage)
	r.POST("/posts/:id", PostsCreateHandler)
	http.ListenAndServe(":3000", r)
}

func GetIndexPage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fp := path.Join("public", "index.html")
	tmpl, _ := template.ParseFiles(fp)
	tmpl.Execute(w, nil)
}

func PostsCreateHandler(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// This can be used to get the parameters from the URL...

	id := p.ByName("id")
	rw.Write([]byte(r.FormValue("userName") + " Value: " + id))
}
