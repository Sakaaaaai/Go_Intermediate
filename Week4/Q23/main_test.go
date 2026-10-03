package main

import (
	"io"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct {
		url      string
		expected string
	}{
		{"http://example.com/", "Hello, Guest!"},
		{"http://example.com/?name=Gopher", "Hello, Gopher!"},
	}

	for _, tt := range tests {
		req := httptest.NewRequest("GET", tt.url, nil)
		w := httptest.NewRecorder()
		
		helloHandler(w, req)
		
		resp := w.Result()
		body, _ := io.ReadAll(resp.Body)
		
		if string(body) != tt.expected {
			t.Errorf("GET %s: expected %q, got %q", tt.url, tt.expected, string(body))
		}
	}
}
