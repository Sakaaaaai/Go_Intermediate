package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 1. リクエストURLのクエリパラメータから "name" を取得してください
	// ヒント: r.URL.Query().Get("キー名") を使います
	name := ___

	// 2. name が空かどうかで分岐し、空なら "Hello, Guest!"、値があれば "Hello, その名前!" を出力してください
	// ヒント: 空文字 "" かどうかを判定し、前問の fmt.Fprint や fmt.Fprintf を使って w に書き込みます
	if ___ {
		___
	} else {
		___
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	fmt.Println("Server is running on http://localhost:8080")
	// テスト用にブラウザで http://localhost:8080/?name=Gopher にアクセスしてみましょう
	http.ListenAndServe(":8080", nil)
}
