package main

import (
	"fmt"
	"time"
)

func Count(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go Count("A")
	Count("B")

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main finished")
}
