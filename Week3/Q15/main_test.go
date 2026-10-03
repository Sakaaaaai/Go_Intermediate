package main

import (
	"os"
	"testing"
)

func TestFileIO(t *testing.T) {
	filename := "test_temp.txt"
	expected := "This is a test string for file IO."

	// 書き込みのテスト
	err := WriteFile(filename, expected)
	if err != nil {
		t.Fatalf("WriteFile failed: %v", err)
	}

	// 読み込みのテスト
	got, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}

	if got != expected {
		t.Errorf("ReadFile() = %q; want %q", got, expected)
	}

	// クリーンアップ
	os.Remove(filename)
}
