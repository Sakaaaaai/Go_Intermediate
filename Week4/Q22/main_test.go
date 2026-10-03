package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// サーバーを直接起動するとテストが難しいため、httptestパッケージを利用します
func TestHTTPServerBasic(t *testing.T) {
	// 期待する動作を行うハンドラを定義
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome to Go Web Server!")
	}

	req := httptest.NewRequest("GET", "http://example.com/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.Status)
	}

	expected := "Welcome to Go Web Server!"
	if string(body) != expected {
		t.Errorf("Expected body %q, got %q", expected, string(body))
	}
}
