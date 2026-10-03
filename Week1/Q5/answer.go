package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("『%s』(著: %s)", b.Title, b.Author)
}

func main() {
	book := Book{
		Title:  "Go言語入門",
		Author: "Go太郎",
	}

	fmt.Println(book)
}
