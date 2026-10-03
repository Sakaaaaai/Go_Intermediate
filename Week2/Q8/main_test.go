package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestGoroutineOutput(t *testing.T) {
	// 実際の出力内容は非同期のため順序が保証されないが、
	// A: 1~3 と B: 1~3 が出力されることを確認する簡易テスト
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "A: 1") || !strings.Contains(output, "A: 3") {
		t.Errorf("Expected output to contain Goroutine A outputs, got: \n%s", output)
	}
	if !strings.Contains(output, "B: 1") || !strings.Contains(output, "B: 3") {
		t.Errorf("Expected output to contain normal B outputs, got: \n%s", output)
	}
	if !strings.Contains(output, "Main finished") {
		t.Errorf("Expected Main finished message")
	}
}
