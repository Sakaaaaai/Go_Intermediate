package main

import "testing"

func BenchmarkLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumLoop(100)
	}
}

func BenchmarkFormula(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumFormula(100)
	}
}
