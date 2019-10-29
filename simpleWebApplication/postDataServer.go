package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/markdown", GenerateMarkdown)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":3000", nil)
}

func GenerateMarkdown(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello: " + r.FormValue("userPassword")))
}
