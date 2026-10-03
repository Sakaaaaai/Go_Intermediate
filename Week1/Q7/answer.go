package main

import (
	"errors"
	"fmt"
)

var ErrPermissionDenied = errors.New("permission denied")

func DoSomething() error {
	return fmt.Errorf("operation failed: %w", ErrPermissionDenied)
}

func main() {
	err := DoSomething()

	if errors.Is(err, ErrPermissionDenied) {
		fmt.Println("アクセスが拒否されたため、処理に失敗しました。")
	} else {
		fmt.Println("その他のエラー:", err)
	}
}
