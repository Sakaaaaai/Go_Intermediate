package main

import (
	"fmt"
)

func FormatValue(v any) string {
	return fmt.Sprintf("Value: %v", v)
}

func main() {
	fmt.Println(FormatValue(42))
	fmt.Println(FormatValue("Hello"))
	fmt.Println(FormatValue(3.14))
	fmt.Println(FormatValue(true))
}
