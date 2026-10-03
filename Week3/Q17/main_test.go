package main

import (
	"testing"
	"time"
)

func TestTimeFormatParse(t *testing.T) {
	// Parse 側のテスト
	parsed, err := ParseTime("2024-05-15")
	if err != nil {
		t.Fatalf("ParseTime failed: %v", err)
	}

	if parsed.Year() != 2024 || parsed.Month() != time.May || parsed.Day() != 15 {
		t.Errorf("ParseTime returned incorrect date: %v", parsed)
	}

	// Format 側のテスト
	// parsed は 2024-05-15 00:00:00
	formatted := FormatTime(parsed)
	expected := "2024/05/15 00:00:00"
	if formatted != expected {
		t.Errorf("FormatTime() = %q; want %q", formatted, expected)
	}
}
