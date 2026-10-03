package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setupRouter はテストのため、muxを生成して返す関数として定義し直します
func setupRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "pong")
	})
	mux.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Go Server Setup")
	})
	return mux
}

func TestRouting(t *testing.T) {
	mux := setupRouter()

	tests := []struct {
		path     string
		expected string
	}{
		{"/ping", "pong"},
		{"/info", "Go Server Setup"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", "http://example.com"+tt.path, nil)
		w := httptest.NewRecorder()
		
		mux.ServeHTTP(w, req)
		
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		
		if string(body) != tt.expected {
			t.Errorf("GET %s expected %q, got %q", tt.path, tt.expected, string(body))
		}
	}
}
