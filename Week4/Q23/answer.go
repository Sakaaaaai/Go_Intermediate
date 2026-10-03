package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello, Guest!")
	} else {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
