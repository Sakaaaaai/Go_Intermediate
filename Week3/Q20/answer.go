package main

import "testing"

func TestIsEven(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, false},
		{0, true},
		{-4, true},
		{-5, false},
	}

	for _, tt := range tests {
		got := IsEven(tt.input)
		if got != tt.expected {
			t.Errorf("IsEven(%d) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}
