package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

type BookRepository interface {
	Save(id string, book Book)
	Get(id string) (Book, bool)
}

type MemoryRepo struct {
	data map[string]Book
}

func (m *MemoryRepo) Save(id string, book Book) {
	m.data[id] = book
}

func (m *MemoryRepo) Get(id string) (Book, bool) {
	book, ok := m.data[id]
	return book, ok
}

func main() {
	var repo BookRepository = &MemoryRepo{data: make(map[string]Book)}

	repo.Save("1", Book{Title: "Go入門", Author: "Gopher"})

	if b, ok := repo.Get("1"); ok {
		fmt.Printf("取得成功: %+v\n", b)
	} else {
		fmt.Println("見つかりません")
	}
}
