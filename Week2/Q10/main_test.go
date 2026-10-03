package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestChannelBasic(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	expected := "Received: Hello from Goroutine!"
	if !strings.Contains(output, expected) {
		t.Errorf("Expected output to contain %q, but got: %q", expected, output)
	}
}
