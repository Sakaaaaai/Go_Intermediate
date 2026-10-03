package main

import "testing"

func TestMemoryRepo(t *testing.T) {
	var repo BookRepository = &MemoryRepo{data: make(map[string]Book)}
	
	book := Book{Title: "テスト本", Author: "著者"}
	repo.Save("100", book)

	got, ok := repo.Get("100")
	if !ok {
		t.Fatalf("Expected book to be found, but got ok=false")
	}
	
	if got.Title != "テスト本" || got.Author != "著者" {
		t.Errorf("Expected %+v, got %+v", book, got)
	}

	_, ok = repo.Get("999")
	if ok {
		t.Fatalf("Expected book not to be found for ID 999, but got ok=true")
	}
}
