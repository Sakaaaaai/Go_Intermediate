package main

import (
	"fmt"
	"io"
	"net/http"
)

func FetchWebsite(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {
	url := "https://httpbin.org/get"
	fmt.Println("Fetching:", url)

	content, err := FetchWebsite(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:")
	fmt.Println(content)
}
