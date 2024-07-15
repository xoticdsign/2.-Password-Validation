package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("./static"))
	http.Handle("/static", http.StripPrefix("/static", h))

	http.HandleFunc("/", MainPageHandler)
	http.HandleFunc("/authorization", AuthorizationHandler)
	http.HandleFunc("/profile", ProfileHandler)

	fmt.Println("Server -> Server started at :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server -> Server failed to start!")
	}
}

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "authorization.html")
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "profile.html")
}
