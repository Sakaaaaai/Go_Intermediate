package main

import (
	"fmt"
	"testing"
)

func TestDescribe(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{100, "整数です: 100"},
		{-5, "整数です: -5"},
		{"Go言語", "文字列です: Go言語"},
		{true, "真偽値です: true"},
		{false, "真偽値です: false"},
		{3.14, "未知の型です"},
		{[]int{1, 2}, "未知の型です"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			if got := Describe(tt.input); got != tt.expected {
				t.Errorf("Describe(%v) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
