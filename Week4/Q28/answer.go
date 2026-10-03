package main

import (
	"errors"
	"fmt"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository interface {
	GetAll() []Book
	GetByID(id string) (Book, error)
	Save(book Book) error
	Delete(id string) error
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

func (m *MemoryRepo) GetByID(id string) (Book, error) {
	book, ok := m.data[id]
	if !ok {
		return Book{}, errors.New("book not found")
	}
	return book, nil
}

func (m *MemoryRepo) Save(book Book) error {
	m.data[book.ID] = book
	return nil
}

func (m *MemoryRepo) Delete(id string) error {
	if _, ok := m.data[id]; !ok {
		return errors.New("book not found")
	}
	delete(m.data, id)
	return nil
}

func main() {
	var repo BookRepository = &MemoryRepo{data: make(map[string]Book)}
	repo.Save(Book{ID: "1", Title: "Go API", Author: "Gopher"})
	fmt.Println("保存されている本:", repo.GetAll())
}
