package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookHandler(t *testing.T) {
	repo := &MemoryRepo{data: make(map[string]Book)}
	repo.data["1"] = Book{ID: "1", Title: "テスト本", Author: "テスト著者"}
	h := &Handler{repo: repo}

	// 1. 正常なGETのテスト
	reqGet := httptest.NewRequest("GET", "/books/1", nil)
	wGet := httptest.NewRecorder()
	h.BookHandler(wGet, reqGet)

	if wGet.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %v", wGet.Code)
	}

	var book Book
	json.NewDecoder(wGet.Body).Decode(&book)
	if book.Title != "テスト本" {
		t.Errorf("GET response is incorrect: %+v", book)
	}

	// 2. 存在しないIDのGETのテスト
	reqGet404 := httptest.NewRequest("GET", "/books/999", nil)
	wGet404 := httptest.NewRecorder()
	h.BookHandler(wGet404, reqGet404)

	if wGet404.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found, got %v", wGet404.Code)
	}

	// 3. 正常なDELETEのテスト
	reqDel := httptest.NewRequest("DELETE", "/books/1", nil)
	wDel := httptest.NewRecorder()
	h.BookHandler(wDel, reqDel)

	if wDel.Code != http.StatusNoContent {
		t.Errorf("Expected status 204 No Content, got %v", wDel.Code)
	}

	// 削除されたか確認
	if _, ok := repo.data["1"]; ok {
		t.Errorf("Book was not deleted from repo")
	}

	// 4. 存在しないIDのDELETEのテスト
	reqDel404 := httptest.NewRequest("DELETE", "/books/1", nil) // すでに削除済み
	wDel404 := httptest.NewRecorder()
	h.BookHandler(wDel404, reqDel404)

	if wDel404.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 Not Found for missing book, got %v", wDel404.Code)
	}
}
