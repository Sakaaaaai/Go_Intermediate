package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestSelectTimeout(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	start := time.Now()
	main()
	elapsed := time.Since(start)

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "Timeout!") {
		t.Errorf("Expected output to contain 'Timeout!', got: %q", output)
	}
	
	if elapsed >= 2*time.Second {
		t.Errorf("Test took %v, but it should have timed out at 1 second", elapsed)
	}
}
