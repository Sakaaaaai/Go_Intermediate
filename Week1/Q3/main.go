package main

import "fmt"

// Describe は渡された値の型に応じて異なるメッセージを返します
func Describe(v any) string {
	// ヒント: Type Switch を使って型を判定してください (`switch val := v.(type)`)
	// ヒント: case で int, string, bool を判定し、最後に default で "未知の型です" を返します
	switch ___ {
	case ___:
		return fmt.Sprintf("整数です: %d", ___)
	case ___:
		return fmt.Sprintf("文字列です: %s", ___)
	case ___:
		return fmt.Sprintf("真偽値です: %t", ___)
	___:
		return "未知の型です"
	}
}

func main() {
	fmt.Println(Describe(100))
	fmt.Println(Describe("Go言語"))
	fmt.Println(Describe(true))
	fmt.Println(Describe(3.14)) // 未知の型
}
