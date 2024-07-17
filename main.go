package main

import (
	"log"
	"net/http"
)

var users = map[string]string{"user": "admin", "admin": "123"}

func main() {

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/", http.StripPrefix("/", fs))

	http.HandleFunc("/authorization", AuthorizationHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/profile", ProfileHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal()
	}
}

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/authorization.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Wrong method", http.StatusMethodNotAllowed)
	}
	usernamePOST := r.FormValue("username")
	passwordPOST := r.FormValue("password")

	for username, password := range users {
		if username == usernamePOST && password == passwordPOST {
			http.Redirect(w, r, "/profile", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/authorization", http.StatusSeeOther)
		}
	}
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/profile.html")
}
