package main

import (
	"fmt"
	"net/http"
)

// 1. Logger はミドルウェアとして機能します
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 2. 処理の前にアクセスされたパス（URLのPath）をログに出力してください
		// ヒント: r.URL.Path を使ってアクセスパスを取得します
		fmt.Println("Requested:", ___)

		// 3. 次のハンドラー（引数で受け取った next）を呼び出してください
		// ヒント: next に w と r を渡して実行します
		___
	}
}

// 本来の処理を行うハンドラー
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello with Logger!")
}

func main() {
	// 4. helloHandler を Logger ミドルウェアで包んで登録してください
	// ヒント: Logger 関数に helloHandler を渡し、その戻り値を引数にします
	http.HandleFunc("/", ___)
	
	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", nil)
}
