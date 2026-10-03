package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestJSONAPI(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/api", nil)
	w := httptest.NewRecorder()
	
	apiHandler(w, req)
	
	resp := w.Result()
	
	// Content-Type ヘッダーのチェック
	if contentType := resp.Header.Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %q", contentType)
	}

	// JSON の構造をデコードしてチェック
	var result Response
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode JSON response: %v", err)
	}
	
	if result.Status != "success" || result.Message != "Go JSON API rocks!" {
		t.Errorf("Unexpected JSON data: %+v", result)
	}
}
