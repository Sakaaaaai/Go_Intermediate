package main

import "fmt"

// 1からnまで順番に足していく関数
func SumLoop(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// 数学の公式（等差数列の和）を使って1発で計算する関数
func SumFormula(n int) int {
	return n * (n + 1) / 2
}

func main() {
	n := 100
	fmt.Println("SumLoop:", SumLoop(n))
	fmt.Println("SumFormula:", SumFormula(n))
}
