package main

import (
	"fmt"
	"net/http"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Requested:", r.URL.Path)
		next(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello with Logger!")
}

func main() {
	http.HandleFunc("/", Logger(helloHandler))
	
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
