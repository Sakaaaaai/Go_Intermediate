package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository interface {
	GetAll() []Book
	Save(book Book) error
}

type MemoryRepo struct {
	data map[string]Book
}

func (m *MemoryRepo) GetAll() []Book {
	books := make([]Book, 0, len(m.data))
	for _, b := range m.data {
		books = append(books, b)
	}
	return books
}

func (m *MemoryRepo) Save(book Book) error {
	m.data[book.ID] = book
	return nil
}

type Handler struct {
	repo BookRepository
}

func (h *Handler) BooksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		books := h.repo.GetAll()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var newBook Book
		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		
		h.repo.Save(newBook)
		w.WriteHeader(http.StatusCreated)
		
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	repo := &MemoryRepo{data: make(map[string]Book)}
	repo.Save(Book{ID: "1", Title: "Go入門", Author: "Gopher"})

	h := &Handler{repo: repo}
	http.HandleFunc("/books", h.BooksHandler)
	http.ListenAndServe(":8080", nil)
}
