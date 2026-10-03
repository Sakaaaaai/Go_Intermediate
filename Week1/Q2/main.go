package main

import (
	"fmt"
)

// FormatValue はどんな型でも受け取れるようにしてください。
// 戻り値は "Value: " の後に渡された値を続けた文字列です。
// ヒント: どんな型でも受け取れるようにするには空インターフェース (any) を使います
// ヒント: 文字列を組み立てて返すには fmt.Sprintf を使い汎用フォーマット %v を活用します
func FormatValue(___) ___ {
	return ___
}

func main() {
	// いろいろな型を渡してみる
	fmt.Println(FormatValue(42))
	fmt.Println(FormatValue("Hello"))
	fmt.Println(FormatValue(3.14))
	fmt.Println(FormatValue(true))
}
