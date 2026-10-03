package main

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	// 1. Content-Type ヘッダーを application/json に設定してください
	// ヒント: w.Header().Set("キー", "値") を使います
	___

	// データを準備
	res := Response{
		Status:  "success",
		Message: "Go JSON API rocks!",
	}

	// 2. Decoder ではなく Encoder を使って、直接 w (レスポンスWriter) に JSON をエンコードして書き込んでください
	// ヒント: json.NewEncoder(w) を呼び出し、続けて Encode(res) メソッドを呼び出します
	___
}

func main() {
	http.HandleFunc("/api", apiHandler)
	http.ListenAndServe(":8080", nil)
}
