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
	// 1. パスのプレフィックス "/books/" を取り除き、残りの文字列を ID として取得してください
	// ヒント: strings.TrimPrefix と r.URL.Path を使います
	id := ___
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	// 2. メソッドによって処理を分岐してください
	// ヒント: r.Method で取得します
	switch ___ {
	case ___: // GET の場合の処理
		// ヒント: h.repo.GetByID を呼び出します
		book, err := ___
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		// ヒント: json.NewEncoder(w).Encode() を使って書き込みます
		___

	case ___: // DELETE の場合の処理
		// ヒント: h.repo.Delete を呼び出します
		err := ___
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		// 3. データなし成功のステータス 204 No Content を設定して返してください
		// ヒント: w.WriteHeader と http.StatusNoContent を使います
		___

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
