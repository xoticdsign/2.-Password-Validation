package main

import (
	"fmt"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("static/"))
	http.Handle("/", http.StripPrefix("/", h))

	fmt.Println("Server -> Server started at :8090")

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Println("Server -> Server failed to start!")
	}
}
