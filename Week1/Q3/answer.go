package main

import "fmt"

func Describe(v any) string {
	switch val := v.(type) {
	case int:
		return fmt.Sprintf("整数です: %d", val)
	case string:
		return fmt.Sprintf("文字列です: %s", val)
	case bool:
		return fmt.Sprintf("真偽値です: %t", val)
	default:
		return "未知の型です"
	}
}

func main() {
	fmt.Println(Describe(100))
	fmt.Println(Describe("Go言語"))
	fmt.Println(Describe(true))
	fmt.Println(Describe(3.14))
}
