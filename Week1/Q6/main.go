package main

import (
	"errors"
	"fmt"
	"os"
)

// ReadFile はわざと存在しないファイルを開き、エラーをラップして返します
func ReadFile() error {
	_, err := os.Open("dummy.txt")
	if err != nil {
		// ヒント: fmt.Errorf と %w を使って元のエラーをラップしてください
		// メッセージは "failed to open file: <元のエラー>" にします
		return ___
	}
	return nil
}

func main() {
	err := ReadFile()
	if err != nil {
		fmt.Println("ラップされたエラー:", err)

		// ヒント: errors.Unwrap() を使って err から元のエラーを取り出してください
		originalErr := ___
		fmt.Println("元のエラー:", originalErr)
	}
}
