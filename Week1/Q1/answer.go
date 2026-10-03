package main

import "fmt"

// Speaker インターフェースの定義
type Speaker interface {
	Speak() string
}

// Dog 構造体の定義とメソッド実装
type Dog struct{}

func (d Dog) Speak() string {
	return "ワンワン"
}

// Cat 構造体の定義とメソッド実装
type Cat struct{}

func (c Cat) Speak() string {
	return "ニャー"
}

// Speaker インターフェースを受け取る関数
func SayHello(s Speaker) string {
	return s.Speak()
}

func main() {
	d := Dog{}
	c := Cat{}

	fmt.Println("Dog:", SayHello(d))
	fmt.Println("Cat:", SayHello(c))
}
