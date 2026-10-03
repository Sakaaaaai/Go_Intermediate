package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello from Goroutine!"
	}()

	msg := <-ch

	fmt.Println("Received:", msg)
}
