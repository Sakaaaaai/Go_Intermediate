package main

import "fmt"

func main() {
	ch := make(chan string, 2)

	ch <- "Apple"
	ch <- "Banana"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
