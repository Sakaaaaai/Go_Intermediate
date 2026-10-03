package main

import (
	"sync"
	"testing"
)

func TestMutexCounter(t *testing.T) {
	c := Counter{counts: make(map[string]int)}
	
	var wg sync.WaitGroup
	// 同時実行テスト（-race フラグでテストすると競合を検知可能）
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("testKey")
		}()
	}
	wg.Wait()

	if c.counts["testKey"] != 1000 {
		t.Errorf("Expected count to be 1000, got %d", c.counts["testKey"])
	}
}
