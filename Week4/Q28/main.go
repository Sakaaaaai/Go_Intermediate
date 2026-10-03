package main

import (
	"errors"
	"fmt"
)

// 1. JSONタグ付きの Book 構造体を定義
type Book struct {
	ID     string `json:"___"`
	Title  string `json:"___"`
	Author string `json:"___"`
}

// 2. 4つのCRUDメソッドを持つ BookRepository インターフェースを定義
type BookRepository interface {
	GetAll() []___
	GetByID(id string) (Book, error)
	___(book Book) error
	___(id string) error
}

// 3. MemoryRepoの実装
type MemoryRepo struct {
	data map[___]___
}

func (m *MemoryRepo) GetAll() []Book {
	// ヒント: nil ではなく空スライスを返すように初期化する
	books := ___
	for _, b := range m.data {
		books = append(books, b)
	}
	return books
}

func (m *MemoryRepo) GetByID(id string) (Book, error) {
	// ヒント: マップから id をキーに検索する
	book, ok := ___
	if !ok {
		return Book{}, errors.New("book not found")
	}
	return book, nil
}

func (m *MemoryRepo) Save(book Book) error {
	// ヒント: ID重複を気にせずマップに保存・上書きする
	___ = book
	return nil
}

func (m *MemoryRepo) Delete(id string) error {
	if _, ok := m.data[id]; !ok {
		return errors.New("book not found")
	}
	// ヒント: マップからの削除に delete 関数を使う
	___
	return nil
}

func main() {
	var repo BookRepository = &MemoryRepo{data: make(map[string]Book)}
	repo.Save(Book{ID: "1", Title: "Go API", Author: "Gopher"})
	fmt.Println("保存されている本:", repo.GetAll())
}
