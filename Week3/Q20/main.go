package main

import "fmt"

// IsEven は与えられた整数が偶数であれば true を返します
func IsEven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println("4 is even?", IsEven(4))
	fmt.Println("5 is even?", IsEven(5))
}
