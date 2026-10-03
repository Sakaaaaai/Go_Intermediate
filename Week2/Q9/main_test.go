package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	// worker関数自体のテスト
	wg.Add(1)
	start := time.Now()
	go worker(1, &wg)
	wg.Wait()
	elapsed := time.Since(start)

	if elapsed < 100*time.Millisecond {
		t.Errorf("worker test finished too quickly, WaitGroup Wait might not be working: %v", elapsed)
	}

	// mainの出力テスト
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "All workers finished") {
		t.Errorf("Expected 'All workers finished' in output")
	}
	
	for i := 1; i <= 3; i++ {
		if !strings.Contains(output, fmt.Sprintf("Worker %d done", i)) {
			t.Errorf("Expected 'Worker %d done' in output", i)
		}
	}
}
