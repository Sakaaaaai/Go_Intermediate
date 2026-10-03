package main

import "testing"

func TestBookRepository(t *testing.T) {
	var repo BookRepository = &MemoryRepo{data: make(map[string]Book)}
	
	book1 := Book{ID: "1", Title: "本1", Author: "著者1"}
	book2 := Book{ID: "2", Title: "本2", Author: "著者2"}
	
	// Save
	repo.Save(book1)
	repo.Save(book2)
	
	// GetAll
	all := repo.GetAll()
	if len(all) != 2 {
		t.Errorf("Expected 2 books, got %d", len(all))
	}
	
	// GetByID
	got, err := repo.GetByID("1")
	if err != nil || got.Title != "本1" {
		t.Errorf("Failed to get correct book: %v", err)
	}
	
	// Delete
	err = repo.Delete("1")
	if err != nil {
		t.Errorf("Failed to delete book: %v", err)
	}
	
	allAfterDelete := repo.GetAll()
	if len(allAfterDelete) != 1 {
		t.Errorf("Expected 1 book after delete, got %d", len(allAfterDelete))
	}
}
