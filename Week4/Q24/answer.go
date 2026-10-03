package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "pong")
	})

	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Go Server Setup")
	})

	fmt.Println("Server starts at :8080")
	http.ListenAndServe(":8080", mux)
}
