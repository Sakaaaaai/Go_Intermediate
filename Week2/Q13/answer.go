package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu     sync.Mutex
	counts map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.counts[key]++
}

func main() {
	c := Counter{counts: make(map[string]int)}

	var wg sync.WaitGroup
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
