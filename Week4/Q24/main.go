package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. 新しいマルチプレクサ（ルーター）を作成してください
	// ヒント: http.NewServeMux() を使います
	mux := ___

	// 2. "/ping" エンドポイントのハンドラーを作成してください
	// ヒント: mux.HandleFunc を使って、レスポンスとして文字列 "pong" を返すようにします
	___("/ping", func(w http.ResponseWriter, r *http.Request) {
		___
	})

	// 3. "/info" エンドポイントのハンドラーを作成してください
	// ヒント: mux.HandleFunc を使って、レスポンスとして文字列 "Go Server Setup" を返すようにします
	___("/info", func(w http.ResponseWriter, r *http.Request) {
		___
	})

	fmt.Println("Server starts at :8080")
	// 4. サーバーを起動してください
	// ヒント: http.ListenAndServe(":8080", ハンドラー) の第二引数に作成した mux を渡します
	http.ListenAndServe(":8080", ___)
}
