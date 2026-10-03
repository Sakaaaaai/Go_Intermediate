package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Process finished"
	}()

	select {
	case msg := <-ch:
		fmt.Println("Success:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout!")
	}
}
