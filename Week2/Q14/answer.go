package main

import (
	"context"
	"fmt"
	"time"
)

func slowProcess(ctx context.Context, done chan struct{}) {
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("処理が正常に完了しました")
	case <-ctx.Done():
		fmt.Println("処理が中断されました:", ctx.Err())
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan struct{})
	fmt.Println("重い処理を開始します...")
	slowProcess(ctx, done)
}
