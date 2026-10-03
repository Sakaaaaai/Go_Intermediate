package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	// 1. sync.Mutex を mu という名前で追加してください
	// ヒント: 変数名 型名 の順に書きます
	___ ___
	counts map[string]int
}

// Inc は指定されたキーのカウントを 1 増やします
func (c *Counter) Inc(key string) {
	// 2. マップを更新する前に排他ロックをかけてください
	// ヒント: mu のロック用メソッドを呼びます
	___

	// 3. 処理が終わったら必ずアンロックされるようにしてください
	// ヒント: defer を使って mu のアンロック用メソッドを呼びます
	___

	c.counts[key]++
}

func main() {
	// カウンターの初期化
	c := Counter{counts: make(map[string]int)}

	var wg sync.WaitGroup
	// 100個のゴルーチンから同時に書き込む
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("apple")
		}()
	}

	wg.Wait()
	fmt.Println("apple count:", c.counts["apple"])
}
