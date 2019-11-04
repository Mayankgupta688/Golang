package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w htt.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World..."))
	}

	http.ListenAndServe(":4000", nil)
}