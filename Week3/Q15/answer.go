package main

import (
	"fmt"
	"os"
)

func WriteFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func ReadFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func main() {
	filename := "test.txt"
	content := "Hello Go!"

	err := WriteFile(filename, content)
	if err != nil {
		fmt.Println("書き込みエラー:", err)
		return
	}

	readContent, err := ReadFile(filename)
	if err != nil {
		fmt.Println("読み込みエラー:", err)
		return
	}

	fmt.Println("ファイルの内容:", readContent)
}
