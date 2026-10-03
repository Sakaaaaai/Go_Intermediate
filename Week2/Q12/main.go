package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	// 2秒かかる重い処理（ゴルーチン）
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Process finished"
	}()

	// 1. select文を使ってチャネルの受信またはタイムアウトを同時に待ち受けます
	// ヒント: select { ... } 構文を使います
	___ {
	// 2. ch から受信した場合
	// ヒント: case msg := <-チャネル: を使います
	___:
		fmt.Println("Success:", msg)

	// 3. time.After で 1秒 のタイムアウトを設定してください
	// ヒント: case <-time.After(期間): を使います
	___:
		fmt.Println("Timeout!")
	}
}
