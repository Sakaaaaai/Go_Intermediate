package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository interface {
	GetByID(id string) (Book, error)
	Delete(id string) error
}

type MemoryRepo struct {
	data map[string]Book
}

func (m *MemoryRepo) GetByID(id string) (Book, error) {
	book, ok := m.data[id]
	if !ok {
		return Book{}, errors.New("not found")
	}
	return book, nil
}

func (m *MemoryRepo) Delete(id string) error {
	if _, ok := m.data[id]; !ok {
		return errors.New("not found")
	}
	delete(m.data, id)
	return nil
}

type Handler struct {
	repo BookRepository
}

func (h *Handler) BookHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/books/")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		book, err := h.repo.GetByID(id)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)

	case http.MethodDelete:
		err := h.repo.Delete(id)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	repo := &MemoryRepo{data: make(map[string]Book)}
	repo.data["1"] = Book{ID: "1", Title: "Go API", Author: "Gopher"}

	h := &Handler{repo: repo}
	http.HandleFunc("/books/", h.BookHandler)
	http.ListenAndServe(":8080", nil)
}
