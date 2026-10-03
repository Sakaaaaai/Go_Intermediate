package main

import "fmt"

// 1. Book 型の定義
type Book struct {
	Title  string
	Author string
}

// 2. BookRepository インターフェースの定義
// ヒント: Save(id string, book Book) と Get(id string) (Book, bool) の2つのメソッドを持ちます
type BookRepository ___ {
	___
	___
}

// 3. MemoryRepo 構造体の定義
type MemoryRepo struct {
	data map[string]Book
}

// MemoryRepo に Save を実装する（BookRepositoryを満たすため）
// ヒント: 構造体のメソッドとして定義します
func (m *MemoryRepo) ___(id string, book Book) {
	// マップのキーに id を指定して保存してください
	m.data[___] = ___
}

// MemoryRepo に Get を実装する（BookRepositoryを満たすため）
// ヒント: 構造体のメソッドとして定義します。戻り値に注意。
func (m *MemoryRepo) ___(id string) (Book, bool) {
	// マップから取得し、存在するかどうかの bool も一緒に返してください
	book, ok := m.data[___]
	return book, ___
}

func main() {
	// 実装した MemoryRepo を インターフェース型変数に代入
	var repo BookRepository = &MemoryRepo{data: make(map[string]Book)}

	repo.Save("1", Book{Title: "Go入門", Author: "Gopher"})

	// 取得して表示
	if b, ok := repo.Get("1"); ok {
		fmt.Printf("取得成功: %+v\n", b)
	} else {
		fmt.Println("見つかりません")
	}
}
