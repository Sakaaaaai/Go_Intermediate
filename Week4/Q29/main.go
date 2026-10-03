package main

import (
	"encoding/json"
	"net/http"
)

// Q28で作成したものと同じ構造を使います
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

// ハンドラーにリポジトリを持たせるための構造体
type Handler struct {
	repo BookRepository
}

func (h *Handler) BooksHandler(w http.ResponseWriter, r *http.Request) {
	// 1. HTTPメソッドによって処理を分けてください
	// ヒント: r.Method でメソッドを取得し、http.MethodGet などを case に指定します
	switch ___ {
	case ___: // GETリクエストの処理
		// 2. h.repo のメソッドを使って本の一覧を取得し、JSONでレスポンスに書き込んでください
		// ヒント: GetAll() メソッドと json.NewEncoder(w).Encode() を使います
		books := ___
		w.Header().Set("Content-Type", "application/json")
		___

	case ___: // POSTリクエストの処理
		// 3. リクエストボディのJSONを読み取って構造体 newBook にデコードしてください
		// ヒント: json.NewDecoder(r.Body).Decode() を使います
		var newBook Book
		if err := ___(&newBook); err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		
		// 取得した newBook を保存してください
		___
		// 4. HTTPステータスコードに 201 Created を設定して返してください
		// ヒント: w.WriteHeader と http.StatusCreated を使います
		___
		
	default:
		// それ以外のメソッドは許可しない
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	repo := &MemoryRepo{data: make(map[string]Book)}
	// テストデータ
	repo.Save(Book{ID: "1", Title: "Go入門", Author: "Gopher"})

	h := &Handler{repo: repo}
	http.HandleFunc("/books", h.BooksHandler)
	http.ListenAndServe(":8080", nil)
}
