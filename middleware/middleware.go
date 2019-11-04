package main

import (
	"fmt"
	"net/http"
)

var Name = "Mayank"

func main() {
	http.HandleFunc("/login", CheckAuthentication(Login))
	http.HandleFunc("/logout", Logout)
	http.ListenAndServe(":4000", nil)
}

func CheckAuthentication(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Authorization Check")

		if Name == "Mayank" {
			Authorize(w, r)
			return
		} else {
			handler(w, r)
		}

	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are logged In..."))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("You are logged Out..."))
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Anauthorized Use..."))
}
