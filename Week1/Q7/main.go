package main

import (
	"errors"
	"fmt"
)

// 1. ErrPermissionDenied を定義してください
var ErrPermissionDenied = ___.___("permission denied")

// 2. ErrPermissionDenied をラップしたエラーを返す関数
func DoSomething() error {
	// "operation failed: <ErrPermissionDenied>" となるようラップする
	return ___("operation failed: %w", ___)
}

func main() {
	err := DoSomething()

	// 3. errors.Is を使って判定してください
	if ___.___(err, ErrPermissionDenied) {
		fmt.Println("アクセスが拒否されたため、処理に失敗しました。")
	} else {
		fmt.Println("その他のエラー:", err)
	}
}
