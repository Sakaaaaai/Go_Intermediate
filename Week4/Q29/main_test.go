package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBooksHandler(t *testing.T) {
	repo := &MemoryRepo{data: make(map[string]Book)}
	h := &Handler{repo: repo}

	// 1. POSTのテスト（作成）
	postBody := `{"id":"1", "title":"テスト", "author":"テストマン"}`
	reqPost := httptest.NewRequest("POST", "/books", bytes.NewBufferString(postBody))
	reqPost.Header.Set("Content-Type", "application/json")
	wPost := httptest.NewRecorder()
	h.BooksHandler(wPost, reqPost)

	if wPost.Code != http.StatusCreated {
		t.Errorf("Expected status 201 Created, got %v", wPost.Code)
	}

	// 2. GETのテスト（取得内容の確認）
	reqGet := httptest.NewRequest("GET", "/books", nil)
	wGet := httptest.NewRecorder()
	h.BooksHandler(wGet, reqGet)

	if wGet.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", wGet.Code)
	}

	var books []Book
	json.NewDecoder(wGet.Body).Decode(&books)
	
	if len(books) != 1 || books[0].Title != "テスト" {
		t.Errorf("GET response is not correct: %+v", books)
	}

	// 3. 不正なメソッドのテスト
	reqPut := httptest.NewRequest("PUT", "/books", nil)
	wPut := httptest.NewRecorder()
	h.BooksHandler(wPut, reqPut)

	if wPut.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected Method Not Allowed (405), got %v", wPut.Code)
	}
}
