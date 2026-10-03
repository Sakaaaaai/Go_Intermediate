package main

import "fmt"

func main() {
	// 1. string型のチャネルを作成してください
	// ヒント: make(chan 型名) を使います
	ch := ___

	// ゴルーチンを起動
	go func() {
		// 2. チャネル ch に "Hello from Goroutine!" を送信してください
		// ヒント: 送信は `チャネル <- 値` のように書きます
		___
	}()

	// 3. チャネル ch からデータを受信し、変数 msg に代入してください
	// ヒント: 受信は `<-チャネル` のように書きます
	msg := ___

	fmt.Println("Received:", msg)
}
