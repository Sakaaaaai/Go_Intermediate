package main

import (
	"fmt"
	"os"
)

// WriteFile は指定したファイル名でテキストを書き込みます
func WriteFile(filename, content string) error {
	// 1. os.WriteFile を使用してファイルを作成・書き込みしてください（権限は 0644）
	// ヒント: 第1引数にファイル名、第2引数にバイト配列([]byte)に変換した書き込みデータ、第3引数に権限 0644 を渡します
	return ___
}

// ReadFile は指定したファイル名のテキストを読み込みます
func ReadFile(filename string) (string, error) {
	// 2. os.ReadFile を使用してファイルを読み込んでください
	// ヒント: 引数にファイル名を渡し、戻り値はバイト配列とエラーの2つです
	data, err := ___
	if err != nil {
		return "", err
	}
	
	// 3. 読み込んだ data（バイト配列）を文字列に変換して返してください
	// ヒント: string(バイト配列) で変換できます
	return ___, nil
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
