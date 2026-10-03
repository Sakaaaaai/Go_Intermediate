package main

import "fmt"

// Add は a と b を足し算した結果を返します
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("2 + 3 =", Add(2, 3))
}
