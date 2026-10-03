package main

import (
	"fmt"
	"time"
)

// Count は、指定された名前と一緒に 1〜3 までの数字を出力します
func Count(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Printf("%s: %d\n", name, i)
		time.Sleep(100 * time.Millisecond) // 出力結果をわかりやすくするための待機
	}
}

func main() {
	// 1. ゴルーチンとして（非同期で） Count("A") を呼び出してください
	// ヒント: 関数の呼び出しの前に go をつけます
	___

	// 2. 通常の関数として（同期的に） Count("B") を呼び出してください
	___

	// メインのゴルーチンが終了するのを防ぐための待機
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main finished")
}
