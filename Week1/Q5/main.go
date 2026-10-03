package main

import "fmt"

// 1. Book 構造体の定義
type Book struct {
	Title  string
	Author string
}

// 2. fmt.Stringer インターフェースを満たすように String() メソッドを実装してください
// ヒント: メソッド名は String、戻り値は string で fmt.Sprintf を使って "『<Title>』(著: <Author>)" 形式の文字列を返します
func ___ {
	return fmt.Sprintf("『%s』(著: %s)", ___, ___)
}

func main() {
	book := Book{
		Title:  "Go言語入門",
		Author: "Go太郎",
	}

	// fmt.Println に渡すと自動的に String() メソッドが呼ばれます
	fmt.Println(book)
}
