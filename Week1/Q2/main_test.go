package main

import (
	"fmt"
	"testing"
)

func TestFormatValue(t *testing.T) {
	tests := []struct {
		input    any
		expected string
	}{
		{42, "Value: 42"},
		{"Hello", "Value: Hello"},
		{3.14, "Value: 3.14"},
		{true, "Value: true"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			if got := FormatValue(tt.input); got != tt.expected {
				t.Errorf("FormatValue(%v) = %q; want %q", tt.input, got, tt.expected)
			}
		})
	}
}
