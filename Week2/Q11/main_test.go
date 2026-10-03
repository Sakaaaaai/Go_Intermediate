package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBufferedChannel(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Apple") {
		t.Errorf("Expected output to contain 'Apple'")
	}
	if !strings.Contains(output, "Banana") {
		t.Errorf("Expected output to contain 'Banana'")
	}
}
