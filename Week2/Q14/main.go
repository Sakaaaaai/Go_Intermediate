package main

import (
	"context"
	"fmt"
	"time"
)

// slowProcess は非常に時間のかかる処理を行います
func slowProcess(ctx ___, done chan struct{}) {
	// ゴルーチンで重い処理をシミュレート（3秒かかる）
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("処理が正常に完了しました")
	// 2. コンテキストのキャンセル（またはタイムアウト）通知を受け取って中断してください
	// ヒント: ctx のメソッドが返すチャネルからの受信を待ち受けます
	case ___:
		fmt.Println("処理が中断されました:", ctx.Err())
	}
}

func main() {
	// 1. 2秒でタイムアウトするコンテキストを作成してください
	// ヒント: context.WithTimeout と context.Background() を使います
	ctx, cancel := ___
	
	// ヒント: メイン関数終了時にタイマーリソースを解放するため defer を使って cancel を呼びます
	___

	done := make(chan struct{})
	fmt.Println("重い処理を開始します...")
	slowProcess(ctx, done)
}
