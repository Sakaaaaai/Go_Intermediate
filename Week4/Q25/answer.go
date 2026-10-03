package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := Response{
		Status:  "success",
		Message: "Go JSON API rocks!",
	}

	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8080", nil)
}
