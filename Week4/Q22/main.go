package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. ルートパス "/" に対するコールバック処理を登録してください
	// ヒント: http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { ... }) の形式になります
	___("/", func(w ___, r *___) {
		// 2. ブラウザ画面に文字列を出力してください
		// ヒント: fmt.Fprint(w, "Welcome to Go Web Server!") を使います
		___
	})

	fmt.Println("Server is running on http://localhost:8080")
	// 3. ポート 8080 でサーバーを起動してください
	// ヒント: http.ListenAndServe を使います。第一引数はポート指定 ":8080" 、第二引数は nil にします
	err := ___
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
