package main

import (
	"bytes"
	"io"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestLoggerMiddleware(t *testing.T) {
	wrappedHandler := Logger(helloHandler)

	req := httptest.NewRequest("GET", "http://example.com/testpath", nil)
	w := httptest.NewRecorder()

	// 標準出力をキャプチャしてログが出ているか確認
	oldStdout := os.Stdout
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	wrappedHandler(w, req)

	wOut.Close()
	os.Stdout = oldStdout
	var buf bytes.Buffer
	io.Copy(&buf, rOut)
	output := buf.String()

	if !strings.Contains(output, "Requested: /testpath") {
		t.Errorf("Expected log containing 'Requested: /testpath', got %q", output)
	}

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)
	if string(body) != "Hello with Logger!" {
		t.Errorf("Expected response body %q, got %q", "Hello with Logger!", string(body))
	}
}
