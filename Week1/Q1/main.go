package main

import "fmt"

// 1. Speak() string メソッドを持つ Speaker インターフェースを定義してください
// ヒント: インターフェース定義のキーワードは interface { ... }
type Speaker ___

// 2. Dog 構造体を定義してください
type Dog struct{}

// Dog に Speak() メソッドを実装してください（"ワンワン" を返す）
// ヒント: 構造体のメソッド：func (レシーバ 型名) メソッド名() 戻り値の型 { ... }
___ {
	return "ワンワン"
}

// 3. Cat 構造体を定義してください
type Cat struct{}

// Cat に Speak() メソッドを実装してください（"ニャー" を返す）
___ {
	return "ニャー"
}

// 4. Speaker インターフェースを受け取り、Speak() の結果を返す関数
// ヒント: 引数の型にインターフェースを指定し、そのメソッドを呼び出します
func SayHello(s ___) string {
	return ___
}

func main() {
	d := Dog{}
	c := Cat{}

	fmt.Println("Dog:", SayHello(d))
	fmt.Println("Cat:", SayHello(c))
}
