package main

import (
	"errors"
	"fmt"
	"os"
)

func ReadFile() error {
	_, err := os.Open("dummy.txt")
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	return nil
}

func main() {
	err := ReadFile()
	if err != nil {
		fmt.Println("ラップされたエラー:", err)

		originalErr := errors.Unwrap(err)
		fmt.Println("元のエラー:", originalErr)
	}
}
