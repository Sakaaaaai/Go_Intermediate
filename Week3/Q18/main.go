package main

import (
	"fmt"
	"io"
	"net/http"
)

// FetchWebsite は指定したURLにアクセスして、その内容を文字列で返します
func FetchWebsite(url string) (string, error) {
	// 1. http.Get を使用してリクエストを送信してください
	// ヒント: 引数に url を渡し、戻り値はレスポンスとエラーの2つです
	resp, err := ___
	if err != nil {
		return "", err
	}

	// 2. 処理終了後に必ずレスポンスの Body を閉じるようにしてください
	// ヒント: defer を使って resp.Body のクローズ用メソッドを呼び出します
	___

	// 3. io.ReadAll を使ってレスポンスの本体(resp.Body)を読み取ってください
	// ヒント: 戻り値はバイト配列とエラーの2つです
	data, err := ___
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func main() {
	url := "https://httpbin.org/get"
	fmt.Println("Fetching:", url)

	content, err := FetchWebsite(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Response:")
	fmt.Println(content)
}
