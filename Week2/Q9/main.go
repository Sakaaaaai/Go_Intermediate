package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	// 3. 終了時に必ず WaitGroup を減らすようにしてください
	// ヒント: defer を使って確実に呼ばれるようにし、wg のメソッドを呼びます
	___

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Duration(id) * 100 * time.Millisecond) // 擬似的な作業時間
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 1. WaitGroupの作成
	// ヒント: 変数名 wg で対象の型を宣言します
	var wg ___

	for i := 1; i <= 3; i++ {
		// 2. ゴルーチンを立ち上げる前に待機数を増やす
		// ヒント: wg のメソッドを呼びます
		___
		
		// ゴルーチンとして worker を起動
		go worker(i, &wg)
	}

	// 4. 全てのワーカーが終了するまで待機する
	// ヒント: wg の待機用メソッドを呼びます
	___

	fmt.Println("All workers finished")
}
