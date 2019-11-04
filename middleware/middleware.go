package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	Name  = "Mayank"
	key   = []byte("sample-secter-key")
	store = sessions.NewCookieStore(key)
)

func main() {
	http.HandleFunc("/login", CheckAuthentication(Login))
	http.HandleFunc("/logout", Logout)
	http.ListenAndServe(":4000", nil)
}

func CheckAuthentication(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		fmt.Println(session.Values["authenticated"] == false)
		if r.URL.Path != "/login" && session.Values["authenticated"] == false {
			Authorize(w, r)
			return
		} else {
			handler(w, r)
		}

	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = true
	session.Save(r, w)
	w.Write([]byte("You are logged In..."))
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	w.Write([]byte("You are logged Out..."))
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Anauthorized Use..."))
}
